package main

import (
	"context"
	"database/sql"
	"ech_measurements/lib"
	"ech_measurements/postgresql"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/debug"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/miekg/dns"
)

// RateLimiter struct to manage the rate of requests
type RateLimiter struct {
	tokens chan struct{}
}

// NewRateLimiter creates a new RateLimiter
func NewRateLimiter(maxRequests int, duration time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, maxRequests),
	}

	go func() {
		ticker := time.NewTicker(duration / time.Duration(maxRequests))
		defer ticker.Stop()

		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}:
			default:
				// Token channel is full; discard
			}
		}
	}()

	return rl
}

// Wait blocks until a token is available
func (rl *RateLimiter) Wait() {
	<-rl.tokens
}

// processDomain processes a single domain and writes the results to the database.
func processDomain(ctx context.Context, id int, domain string, queriedNS string, timestamp time.Time, testCode string, vantagePoint string, worker string, rl *RateLimiter, db *sql.DB, testDate string, writeQueue chan<- postgresql.DNSResult) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic occurred while processing domain %s: %v\n", domain, r)
			debug.PrintStack() // Print the stack trace
		}
	}()

	fmt.Printf("Worker %d processing domain: %s\n", id, domain)
	rl.Wait()

	// Initialize the result variable
	var result postgresql.DNSResult

	// Call QueryDNSRecords with the correct parameters
	result, err := lib.QueryDNSRecords(domain, queriedNS, func(d string, q string, rType uint16, status *error, network *string, attempts *int, rtt_ns *int64, rCode *string, maxAttempts *int, maxBackoff *time.Duration, dialTimeout *time.Duration, queryTimeout *time.Duration) ([]dns.RR, error) {
		return lib.QueryDNSRecord(d, q, rType, status, network, attempts, rtt_ns, rCode, maxAttempts, maxBackoff, dialTimeout, queryTimeout)
	})
	if err != nil {
		log.Printf("Error querying DNS records for domain %s: %v\n", domain, err)
		return
	}

	result.SCRIPTSTARTTIMESTAMP = timestamp
	result.TestCode = testCode
	result.TESTDATE = testDate
	result.VantagePoint = vantagePoint
	result.WORKER = worker

	// Fetch and populate IP-related information
	fetchAndPopulateIPInfo(&result) // Pass result as a pointer
	// Send the processed result to the write queue
	writeQueue <- result
	// Write results to the database with retry mechanism
	/*if err := postgresql.WriteResultsWithRetry(ctx, db, result); err != nil {
		log.Printf("Error writing result to database for domain %s: %v\n", result.Domain, err)
	}*/
}

// fetchAndPopulateIPInfo fetches and populates IP-related information for A, AAAA, and NS records.
func fetchAndPopulateIPInfo(result *postgresql.DNSResult) {
	// Use a WaitGroup to wait for all IP info fetches to complete
	var wg sync.WaitGroup
	wg.Add(3) // 3 groups of records (A, AAAA, NS)
	fmt.Printf("Fetching IP info and IP for domain: %s\n", result.Domain)
	go func() {
		defer wg.Done()
		result.AASN, result.AASNORG, result.AASNCITY, result.AASNCOUNTRY = fetchIPInfoForRecords(result.AVALUE)
	}()

	go func() {
		defer wg.Done()
		result.AAAAASN, result.AAAAASNORG, result.AAAAASNCITY, result.AAAAASNCOUNTRY = fetchIPInfoForRecords(result.AAAAVALUE)
	}()

	go func() {
		defer wg.Done()
		result.NSASN, result.NSASNORG, result.NSASNCITY, result.NSASNCOUNTRY, result.NSIP, result.NSIPVERSION = fetchIPInfoForNSRecords(result.NSVALUE)
	}()

	wg.Wait()
}

// fetchIPInfoForRecords fetches IP information for a slice of IP addresses.
func fetchIPInfoForRecords(records []string) ([]string, []string, []string, []string) {
	asn := make([]string, 0, len(records))
	org := make([]string, 0, len(records))
	city := make([]string, 0, len(records))
	country := make([]string, 0, len(records))

	for _, record := range records {
		if record != "" {
			ipInfo, err := lib.FetchIPInfo(record) // Consider caching IP info here
			if err != nil {
				log.Printf("Error fetching IP info for %s: %v\n", record, err)
			} else {
				asn = append(asn, ipInfo.ASN)
				org = append(org, ipInfo.Organization)
				city = append(city, ipInfo.City)
				country = append(country, ipInfo.Country)
			}
		}
	}

	return asn, org, city, country
}

// fetchIPInfoForNSRecords fetches IP information for NS records, including resolving their IPs.
func fetchIPInfoForNSRecords(records []string) ([]string, []string, []string, []string, []string, []string) {
	asn := make([]string, 0, len(records))
	org := make([]string, 0, len(records))
	city := make([]string, 0, len(records))
	country := make([]string, 0, len(records))
	ipList := make([]string, 0, len(records))
	ipVersion := make([]string, 0, len(records))

	for _, record := range records {
		// Check if the record is already an IP address
		if len(record) > 0 && record[len(record)-1] == '.' {
			record = record[:len(record)-1]
		}
		if net.ParseIP(record) != nil {
			ipString := record
			ip := net.ParseIP(ipString)
			ipInfo, err := lib.FetchIPInfo(ipString) // Consider caching IP info here
			if err != nil {
				log.Printf("Error fetching IP info for %s: %v\n", ipString, err)
			} else {
				asn = append(asn, ipInfo.ASN)
				org = append(org, ipInfo.Organization)
				city = append(city, ipInfo.City)
				country = append(country, ipInfo.Country)
				ipList = append(ipList, ipString)

				if ip.To4() != nil {
					ipVersion = append(ipVersion, "IPv4")
				} else {
					ipVersion = append(ipVersion, "IPv6")
				}
			}
		} else {
			// If not an IP address, fetch it using FetchIPWithRetry
			fmt.Printf("Fetching IP for domain in else: %s\n", record)
			ips, err := lib.FetchIPWithRetry(record) // Consider caching IPs here
			if err != nil {
				log.Printf("Error fetching IP for domain %s: %v\n", record, err)
				continue
			}

			for _, ip := range ips {
				ipString := fmt.Sprintf("%s", ip)
				if ipString != "" {
					ipInfo, err := lib.FetchIPInfo(ipString) // Consider caching IP info here
					if err != nil {
						log.Printf("Error fetching IP info for %s: %v\n", ipString, err)
					} else {
						asn = append(asn, ipInfo.ASN)
						org = append(org, ipInfo.Organization)
						city = append(city, ipInfo.City)
						country = append(country, ipInfo.Country)
						ipList = append(ipList, ipString)

						if ip.To4() != nil {
							ipVersion = append(ipVersion, "IPv4")
						} else {
							ipVersion = append(ipVersion, "IPv6")
						}
					}
				}
			}
		}
	}

	return asn, org, city, country, ipList, ipVersion
}

func main() {
	// Start pprof for performance profiling
	go func() {
		fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}

	// Read environment variables
	vantagePoint := os.Getenv("VANTAGE_POINT")
	querriedNS := os.Getenv("QUERRIED_NS")
	queryThreats, _ := strconv.Atoi(os.Getenv("QUERYTHREATS"))
	dbThreats, _ := strconv.Atoi(os.Getenv("DB_THREATS"))
	testCode := os.Getenv("TESTCODE")
	testString := os.Getenv("TEST")
	test, _ := strconv.ParseBool(testString)
	variety := os.Getenv("VARIETY") // can be  0-3
	worker := os.Getenv("WORKER")
	retry, _ := strconv.ParseBool(os.Getenv("RETRY"))
	var testDate string

	fmt.Printf("vantagepoint: %s, testcode: %s, test: %t, variety: %s, worker: %s, retry: %t\n", vantagePoint, testCode, test, variety, worker, retry)

	// Read rate limiting configuration from .env
	maxRequests, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_MAX_REQUESTS"))
	rateLimitDuration, _ := time.ParseDuration(os.Getenv("RATE_LIMIT_DURATION"))

	//Get time
	timestamp := time.Now()
	timestampSTRING := timestamp.Format("2006-01-02_15-04-05")

	// Initialize the database connection
	db, err := postgresql.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v\n", err)
	}
	defer db.Close()

	// Start monitoring the connection pool statistics
	go func() {
		for {
			stats := db.Stats()
			fmt.Printf("Connection Pool Stats - MaxOpenConnections: %d, OpenConnections: %d, InUse: %d, Idle: %d\n",
				stats.MaxOpenConnections, stats.OpenConnections, stats.InUse, stats.Idle)
			time.Sleep(30 * time.Second) // Adjust the interval as needed
		}
	}()

	// Get the list of domains to query
	domains, err := getDomainsToQuery(db, test, testCode, variety, retry)
	if err != nil {
		log.Fatalf("Failed to get domains: %v\n", err)
	}
	fmt.Printf("Number of domains: %d\n", len(domains))

	// Generate the test code
	if testCode == "" {
		testCode = fmt.Sprintf("%s_%s_%s", timestampSTRING, vantagePoint, querriedNS)
	}

	if testCode != "" {
		testDate = testCode[:10]
	}

	// Initialize the rate limiter
	rl := NewRateLimiter(maxRequests, rateLimitDuration)

	var wg sync.WaitGroup
	var wgWrite sync.WaitGroup
	ctx := context.Background() // Create a context for cancellation

	// Start worker goroutines
	jobs := make(chan string, len(domains))               // Create a channel for jobs
	writeQueue := make(chan postgresql.DNSResult, 200000) // Adjust buffer size as needed

	for w := 1; w <= queryThreats; w++ {
		wg.Add(1)
		go queryWorker(ctx, w, jobs, &wg, querriedNS, timestamp, testCode, vantagePoint, worker, rl, db, testDate, writeQueue)
	}

	// Start DB write worker goroutines
	for w := 1; w <= dbThreats; w++ {
		wgWrite.Add(1)
		go writeWorker(ctx, w, db, &wgWrite, writeQueue)
	}

	// Dispatch jobs to workers
	for _, domain := range domains {
		jobs <- domain
	}
	close(jobs) // Signal that no more jobs will be sent

	// Wait for all workers to finish
	wg.Wait()
	close(writeQueue)

	// Wait for all MySQL write workers to finish
	wgWrite.Wait()
	fmt.Println("DNS record processing and database writing complete.")
}

// getDomainsToQuery retrieves the list of domains to query based on various conditions.
func getDomainsToQuery(db *sql.DB, test bool, testCode, variety string, retry bool) ([]string, error) {
	if test {
		return []string{
			"defo.ie",
			/*"hpi.de",
			"hcc.nl",
			"xolus.net",
			"mmcert.org.mm",
			"internetsociety.org",
			"cz.nic",
			"dnssec-tools.org",
			"uni-potsdam.de",
			"sdsu.com",
			"example.com",
			"nato.gov",
			"google.com",
			"amazon.com",
			"grosse-junkersdorfer.de",
			"0-g-0.ru",
			"0-www.nhlbi.nih.gov.innopac.up.ac.za",
			"cloudflare.com",
			"mozilla.org",
			"microsoft.com",
			"research.cloudflare.com",
			"GitHub.com",
			"Netflix.com",
			"Wikipedia.org",
			"GitHub.com",
			"tlsa.huque.com",
			"huque.com",
			"sys4.de",
			"cdn.cloudflare.net",
			"www.github.com",
			"www.cloudflare-ech.com",
			"dns.google",
			"www.fastly.com",
			"dnsviz.net",
			"verisignlabs.com",
			"nlnetlabs.nl",
			"ietf.org",*/
		}, nil
	}

	if testCode == "" && !test {
		fmt.Print("Fetching domains from the database...with GetDomains\n")
		return postgresql.GetDomains(db)
	}

	if testCode != "" && variety == "" && retry && !test {
		fmt.Printf("Fetching untested domains from the database...with GetDomainsForReassessmentByTestCode\n")
		return postgresql.GetDomainsForReassessmentByTestCode(db, testCode)
	}

	if testCode != "" && variety != "" && !retry && !test {
		varietyInt, _ := strconv.Atoi(variety)

		fmt.Printf("Fetching untested domains from the database...with GetDomainsByModulo\n")
		return postgresql.GetDomainsByModulo(db, varietyInt, 5)
	}

	if testCode != "" && variety != "" && retry && !test {
		varietyInt, _ := strconv.Atoi(variety)
		fmt.Printf("Fetching domains from the database...with GetUntestedDomainsByModulo\n")
		return postgresql.GetUntestedDomainsByModulo(db, testCode, varietyInt, 5)
	}
	fmt.Print("Fetching domains from the database...with GetDomains no if condition matched\n")
	return postgresql.GetDomains(db)
}

// queryWorker is a worker function that processes domains from the jobs channel.
func queryWorker(ctx context.Context, id int, jobs <-chan string, wg *sync.WaitGroup, queriedNS string, timestamp time.Time, testCode string, vantagePoint string, worker string, rl *RateLimiter, db *sql.DB, testDate string, writeQueue chan<- postgresql.DNSResult) {
	defer wg.Done()
	for domain := range jobs { // Receive domains from the channel
		select {
		case <-ctx.Done():
			return // Exit if the context is canceled
		default:
			processDomain(ctx, id, domain, queriedNS, timestamp, testCode, vantagePoint, worker, rl, db, testDate, writeQueue)
		}
	}
}

func writeWorker(ctx context.Context, id int, database *sql.DB, wg *sync.WaitGroup, writeQueue <-chan postgresql.DNSResult) {
	defer wg.Done()
	for result := range writeQueue {
		// Write each result to the database
		fmt.Printf("Write Worker %d writing result for domain: %s\n", id, result.Domain)
		if err := postgresql.WriteResultsWithRetry(ctx, database, result); err != nil {
			fmt.Printf("Error writing result to database for domain %s: %v\n", result.Domain, err)
		}
	}
}
