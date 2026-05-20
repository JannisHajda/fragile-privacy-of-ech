package postgresql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// DNSResult represents a DNS query result with various record types.
type DNSResult struct {
	Domain                 string
	SCRIPTSTARTTIMESTAMP   time.Time
	DNSQUERYSTARTTIMESTAMP time.Time
	DNSQUERYENDTIMESTAMP   time.Time
	QUERRIEDNS             string
	VantagePoint           string
	ErrorCode              []string
	TestCode               string
	TESTDATE               string
	DomainID               string
	WORKER                 string
	DIALTIMEOUT            time.Duration
	QUERYTIMEOUT           time.Duration
	MAXATTEMPTS            int
	MAXBACKOFF             time.Duration

	ARAW               []string //Changed
	AHOST              []string
	ATTL               []uint32 //Changed
	AVALUE             []string
	AASN               []string
	AASNORG            []string
	AASNCITY           []string
	AASNCOUNTRY        []string
	ARRTYPE            []uint16
	ARDLEN             []uint16
	ASIGNED            bool
	AVERIFIED          bool
	AVERIFICATIONERROR string
	AERRORMSG          string
	ARTT               int64
	AATTEMPTS          int
	ANETWORK           string
	ARRSIGERRORMSG     string
	AVALIDPERIOD       bool
	ARCODE             string

	AAAARAW               []string //Changed
	AAAAHOST              []string
	AAAATTL               []uint32 //Changed
	AAAAVALUE             []string
	AAAAASN               []string
	AAAAASNORG            []string
	AAAAASNCITY           []string
	AAAAASNCOUNTRY        []string
	AAAARRTYPE            []uint16
	AAAARDLEN             []uint16
	AAAASIGNED            bool
	AAAAVERIFIED          bool
	AAAAVERIFICATIONERROR string
	AAAAERRORMSG          string
	AAAARTT               int64
	AAAAATTEMPTS          int
	AAAANETWORK           string
	AAAARRSIGERRORMSG     string
	AAAAVALIDPERIOD       bool
	AAAARCODE             string

	CNAMERAW               []string //Changed
	CNAMEHOST              []string //Changed
	CNAMETTL               []uint32 //Changed
	CNAMEVALUE             []string //Changed
	CNAMERRTYPE            []uint16
	CNAMERDLEN             []uint16
	CNAMESIGNED            bool
	CNAMEVERIFIED          bool
	CNAMEVERIFICATIONERROR string
	CNAMEERRORMSG          string
	CNAMERTT               int64
	CNAMEATTEMPTS          int
	CNAMENETWORK           string
	CNAMERRSIGERRORMSG     string
	CNAMEVALIDPERIOD       bool
	CNAMERCODE             string

	MXRAW               []string //Changed
	MXHOST              []string
	MXTTL               []uint32 //Changed
	MXVALUE             []string
	MXIP                []string
	MXIPVERSION         []string
	MXASN               []string
	MXASNORG            []string
	MXASNCITY           []string
	MXASNCOUNTRY        []string
	MXRRTYPE            []uint16
	MXRDLEN             []uint16
	MXSIGNED            bool
	MXVERIFIED          bool
	MXVERIFICATIONERROR string
	MXERRORMSG          string
	MXRTT               int64
	MXATTEMPTS          int
	MXNETWORK           string
	MXRRSIGERRORMSG     string
	MXVALIDPERIOD       bool
	MXRCODE             string

	NSRAW               []string //Changed
	NSHOST              []string
	NSTTL               []uint32 //Changed
	NSVALUE             []string
	NSIP                []string
	NSIPVERSION         []string
	NSASN               []string
	NSASNORG            []string
	NSASNCITY           []string
	NSASNCOUNTRY        []string
	NSRRTYPE            []uint16
	NSRDLEN             []uint16
	NSSIGNED            bool
	NSVERIFIED          bool
	NSVERIFICATIONERROR string
	NSERRORMSG          string
	NSRTT               int64
	NSATTEMPTS          int
	NSNETWORK           string
	NSRRSIGERRORMSG     string
	NSVALIDPERIOD       bool
	NSRCODE             string

	TXTRAW               []string //Changed
	TXTHOST              []string
	TXTTTL               []uint32 //Changed
	TXTVALUE             []string
	TXTCLASS             []uint16
	TXTRRTYPE            []uint16
	TXTRDLEN             []uint16
	TXTSIGNED            bool
	TXTVERIFIED          bool
	TXTVERIFICATIONERROR string
	TXTERRORMSG          string
	TXTRTT               int64
	TXTATTEMPTS          int
	TXTNETWORK           string
	TXTRRSIGERRORMSG     string
	TXTVALIDPERIOD       bool
	TXTRCODE             string

	SVCBRAW               []string
	SVCBNAME              []string
	SVCBTARGET            []string
	SVCBTTL               []uint32
	SVCBVALUE             []string
	SVCBCLASS             []uint16
	SVCBRDLEN             []uint16
	SVCBRRTYPE            []uint16
	SVCBPRIORITY          []uint16
	SVCBSIGNED            bool
	SVCBVERIFIED          bool
	SVCBVERIFICATIONERROR string
	SVCBERRORMSG          string
	SVCBRTT               int64
	SVCBATTEMPTS          int
	SVCBNETWORK           string
	SVCBRRSIGERRORMSG     string
	SVCBVALIDPERIOD       bool
	SVCBRCODE             string

	HTTPSRAW               string
	HTTPSPRIORITY          uint32
	HTTPSTARGET            string
	HTTPSNAME              string
	HTTPSCLASS             uint16
	HTTPSTTL               uint32
	HTTPSIPV4HINT          string
	HTTPSIPV6HINT          string
	HTTPSECHKEY            string
	HTTPSALPN              string
	HTTPS0HTTP             string
	HTTPSODOHCONFIG        string
	HTTPSMANDATORY         string
	HTTPSPORT              uint32
	HTTPSNODEFAULALPN      string
	HTTPSDOHTTARGET        string
	HTTPSESNIKEYS          string
	HTTPSRDLEN             uint16
	HTTPSRRTYPE            uint16
	HTTPSSIGNED            bool
	HTTPSVERIFIED          bool
	HTTPSVERIFICATIONERROR string
	HTTPSERRORMSG          string
	HTTPSRTT               int64
	HTTPSATTEMPTS          int
	HTTPSNETWORK           string
	HTTPSRRSIGERRORMSG     string
	HTTPSVALIDPERIOD       bool
	HTTPSRCODE             string

	ECHCONFIGB64        []string
	ECHCONFIGRAW        [][]byte
	ECHCONFIGTLSVERSION []string
	ECHCONFIGLENGTH     []string
	ECHCONFIGID         []string
	ECHKEMID            []string
	ECHPUBLICKEY        []string
	ECHMAXNAMELENGTH    []string
	ECHPUBLICNAME       []string
	ECHCIPHERKDFID      []string
	ECHCIPHERAEADID     []string
	ECHEXTENSIONTYPE    []string
	ECHEXTENSIONDATA    [][]byte

	RRSIGRAW               []string
	RRSIGNAME              []string
	RRSIGTTL               []uint32
	RRSIGCLASS             []uint16
	RRSIGRRTYPE            []uint16
	RRSIGRDLEN             []uint16
	RRSIGEXPIRATION        []string
	RRSIGINCEPTION         []string
	RRSIGKEYTAG            []uint16
	RRSIGKEYTAGSTRING      []string
	RRSIGTYPECOVERED       []uint16
	RRSIGTYPECOVEREDSTRING []string
	RRSIGLABELS            []uint8
	RRSIGALGORITHM         []uint8
	RRSIGALGORITHMSTRING   []string
	RRSIGORIGTTL           []uint32
	RRSIGSIGNATURE         []string
	RRSIGSIGNERNAME        []string
	//RRSIGVALIDPERIOD       []string
	RRSIGTIMEVALIDPERIOD   []string
	RRSIGVALIDSIGKEYNAME   []string
	RRSIGSIGNED            bool
	RRSIGVERIFIED          bool
	RRSIGVERIFICATIONERROR string
	RRSIGERRORMSG          string
	RRSIGRTT               int64
	RRSIGATTEMPTS          int
	RRSIGNETWORK           string
	RRSIGRRSIGERRORMSG     string
	RRSIGVALIDPERIOD       bool
	RRSIGRCODE             string

	DNSKEYRAW               []string
	DNSKEYCLASS             []uint16
	DNSKEYRDLEN             []uint16
	DNSKEYRRTYPE            []uint16
	DNSKEYTTL               []uint32
	DNSKEYNAME              []string
	DNSKEYALGORITHM         []uint8
	DNSKEYALGORITHMSTRING   []string
	DNSKEYFLAGS             []uint16
	DNSKEYPUBLICKEY         []string
	DNSKEYKEYTAG            []uint16
	DNSKEYKEYTAGSTRING      []string
	DNSKEYSIGNED            bool
	DNSKEYVERIFIED          bool
	DNSKEYVERIFICATIONERROR string
	DNSKEYERRORMSG          string
	DNSKEYRTT               int64
	DNSKEYATTEMPTS          int
	DNSKEYNETWORK           string
	DNSKEYRRSIGERRORMSG     string
	DNSKEYVALIDPERIOD       bool
	DNSKEYRCODE             string

	CAARAW               []string
	CAACLASS             []uint16
	CAARDLEN             []uint16
	CAARRTYPE            []uint16
	CAATTL               []uint32
	CAANAME              []string
	CAAFLAG              []uint8
	CAATAG               []string
	CAACONTENT           []string
	CAAVALUE             []string
	CAASIGNED            bool
	CAAVERIFIED          bool
	CAAVERIFICATIONERROR string
	CAAERRORMSG          string
	CAARTT               int64
	CAAATTEMPTS          int
	CAANETWORK           string
	CAARRSIGERRORMSG     string
	CAAVALIDPERIOD       bool
	CAARCODE             string

	SOARAW               []string
	SOACLASS             []uint16
	SOARDLEN             []uint16
	SOARRTYPE            []uint16
	SOATTL               []uint32
	SOANAME              []string
	SOAEXPIRE            []uint32
	SOAMBOX              []string
	SOAMINTTL            []uint32
	SOAREFRESH           []uint32
	SOAREFRESHSTRING     []string
	SOANS                []string
	SOARETRY             []uint32
	SOASERIAL            []uint32
	SOASIGNED            bool
	SOAVERIFIED          bool
	SOAVERIFICATIONERROR string
	SOAERRORMSG          string
	SOARTT               int64
	SOAATTEMPTS          int
	SOANETWORK           string
	SOARRSIGERRORMSG     string
	SOAVALIDPERIOD       bool
	SOARCODE             string

	DSRAW               []string
	DSCLASS             []uint16
	DSRDLEN             []uint16
	DSRRTYPE            []uint16
	DSTTL               []uint32
	DSNAME              []string
	DSALGORITHM         []uint8
	DSKEYTAG            []uint16
	DSKEYTAGSTRING      []string
	DSFLAGS             []string
	DSDIGEST            []string
	DSDIGESTTYPE        []uint8
	DSALGORITHMSTRING   []string
	DSSIGNED            bool
	DSVERIFIED          bool
	DSVERIFICATIONERROR string
	DSERRORMSG          string
	DSRTT               int64
	DSATTEMPTS          int
	DSNETWORK           string
	DSRRSIGERRORMSG     string
	DSVALIDPERIOD       bool
	DSRCODE             string

	CDSRAW               []string
	CDSCLASS             []uint16
	CDSRDLEN             []uint16
	CDSRRTYPE            []uint16
	CDSTTL               []uint32
	CDSNAME              []string
	CDSALGORITHM         []uint8
	CDSALGORITHMSTRING   []string
	CDSKEYTAG            []uint16
	CDSKEYTAGSTRING      []string
	CDSSIGNED            bool
	CDSVERIFIED          bool
	CDSVERIFICATIONERROR string
	CDSERRORMSG          string
	CDSRTT               int64
	CDSATTEMPTS          int
	CDSNETWORK           string
	CDSRRSIGERRORMSG     string
	CDSVALIDPERIOD       bool
	CDSRCODE             string

	CDNSKEYRAW               []string
	CDNSKEYCLASS             []uint16
	CDNSKEYRDLEN             []uint16
	CDNSKEYRRTYPE            []uint16
	CDNSKEYTTL               []uint32
	CDNSKEYNAME              []string
	CDNSKEYALGORITHM         []uint8
	CDNSKEYALGORITHMSTRING   []string
	CDNSKEYFLAGS             []uint16
	CDNSKEYPROTOCOL          []uint8
	CDNSKEYPUBLICKEY         []string
	CDNSKEYDNSKEYTOCDNSKEY   []string
	CDNSKEYKEYTAG            []uint16
	CDNSKEYKEYTAGSTRING      []string
	CDNSKEYDNSKEY            []string
	CDNSKEYSIGNED            bool
	CDNSKEYVERIFIED          bool
	CDNSKEYVERIFICATIONERROR string
	CDNSKEYERRORMSG          string
	CDNSKEYRTT               int64
	CDNSKEYATTEMPTS          int
	CDNSKEYNETWORK           string
	CDNSKEYRRSIGERRORMSG     string
	CDNSKEYVALIDPERIOD       bool
	CDNSKEYRCODE             string

	NSECRAW               []string
	NSECCLASS             []uint16
	NSECRDLEN             []uint16
	NSECRRTYPE            []uint16
	NSECTTL               []uint32
	NSECNAME              []string
	NSECTYPEBITMAPSTRING  []string
	NSECNEXTDOMAIN        []string
	NSECSIGNED            bool
	NSECVERIFIED          bool
	NSECVERIFICATIONERROR string
	NSECERRORMSG          string
	NSECRTT               int64
	NSECATTEMPTS          int
	NSECNETWORK           string
	NSECRRSIGERRORMSG     string
	NSECVALIDPERIOD       bool
	NSECRCODE             string

	NSEC3RAW               []string
	NSEC3NAME              []string
	NSEC3CLASS             []uint16
	NSEC3RDLEN             []uint16
	NSEC3RRTYPE            []uint16
	NSEC3TTL               []uint32
	NSEC3HASHLENGTHSTRING  []string
	NSEC3HASHLENGTH        []uint8
	NSEC3NEXTDOMAIN        []string
	NSEC3HASH              []uint8
	NSEC3ITERATIONS        []uint16
	NSEC3SALT              []string
	NSEC3SALTLENGTH        []uint8
	NSEC3TYPEBITMAPSTRING  []string
	NSEC3SIGNED            bool
	NSEC3VERIFIED          bool
	NSEC3VERIFICATIONERROR string
	NSEC3ERRORMSG          string
	NSEC3RTT               int64
	NSEC3ATTEMPTS          int
	NSEC3NETWORK           string
	NSEC3RRSIGERRORMSG     string
	NSEC3VALIDPERIOD       bool
	NSEC3RCODE             string

	NSEC3PARAMRAW               []string
	NSEC3PARAMNAME              []string
	NSEC3PARAMCLASS             []uint16
	NSEC3PARAMRDLEN             []uint16
	NSEC3PARAMRRTYPE            []uint16
	NSEC3PARAMTTL               []uint32
	NSEC3PARAMHASH              []uint8
	NSEC3PARAMITERATIONS        []uint16
	NSEC3PARAMSALT              []string
	NSEC3PARAMSALTLENGTH        []uint8
	NSEC3PARAMTYPEBITMAPSTRING  []string
	NSEC3PARAMSIGNED            bool
	NSEC3PARAMVERIFIED          bool
	NSEC3PARAMVERIFICATIONERROR string
	NSEC3PARAMERRORMSG          string
	NSEC3PARAMRTT               int64
	NSEC3PARAMATTEMPTS          int
	NSEC3PARAMNETWORK           string
	NSEC3PARAMRRSIGERRORMSG     string
	NSEC3PARAMVALIDPERIOD       bool
	NSEC3PARAMRCODE             string

	TLSARAW               []string
	TLSACLASS             []uint16
	TLSARDLEN             []uint16
	TLSARRTYPE            []uint16
	TLSATTL               []uint32
	TLSANAME              []string
	TLSAUSAGE             []uint8
	TLSACERTIFICATE       []string
	TLSAMATCHINGTYPE      []uint8
	TLSASELECTOR          []uint8
	TLSATARGET            []string
	TLSASIGNED            bool
	TLSAVERIFIED          bool
	TLSAVERIFICATIONERROR string
	TLSAERRORMSG          string
	TLSARTT               int64
	TLSAATTEMPTS          int
	TLSANETWORK           string
	TLSARRSIGERRORMSG     string
	TLSAVALIDPERIOD       bool
	TLSARCODE             string

	OPTRAW               []string
	OPTTTL               []uint32
	OPTCLASS             []uint16
	OPTRDLEN             []uint16
	OPTRRTYPE            []uint16
	OPTNAME              []string
	OPTOPTION            []string
	OPTDOMAIN            []string
	OPTDO                []bool
	OPTEXTENDEDRCODE     []int
	OPTVERSION           []uint8
	OPTUDPSIZE           []uint16
	OPTZ                 []uint16
	OPTSIGNED            bool
	OPTVERIFIED          bool
	OPTVERIFICATIONERROR string
	OPTERRORMSG          string
	OPTRTT               int64
	OPTATTEMPTS          int
	OPTNETWORK           string
	OPTRRSIGERRORMSG     string
	OPTVALIDPERIOD       bool
	OPTRCODE             string
}

// Domain struct to hold domain data.
type Domain struct {
	ID   int64
	Name string
}

// InitDB initializes the database connection with connection pooling.
func InitDB() (*sql.DB, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Get database configuration from environment variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	// Read connection pool settings from environment variables
	maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		maxOpenConns = 1 // default value if not set or invalid
	}

	maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		maxIdleConns = 1 // default value if not set or invalid
	}

	connMaxLifetime, err := time.ParseDuration(os.Getenv("DB_CONN_MAX_LIFETIME"))
	if err != nil {
		connMaxLifetime = time.Minute * 7 // default value if not set or invalid
	}

	// Build the DSN (Data Source Name) string for PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)

	// Open the database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Configure connection pooling settings
	db.SetMaxOpenConns(maxOpenConns)       // Maximum number of open connections to the database
	db.SetMaxIdleConns(maxIdleConns)       // Maximum number of idle connections in the pool
	db.SetConnMaxLifetime(connMaxLifetime) // Maximum lifetime of a connection

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	return db, nil
}

func EscapeSpecialChars(str string) string {
	var escaped strings.Builder

	for i := 0; i < len(str); {
		// Decode the next rune (character) in the string
		r, size := utf8.DecodeRuneInString(str[i:])

		switch {
		case r == utf8.RuneError && size == 1: // Ungültiges Zeichen gefunden
			escaped.WriteRune('�') // Platzhalter-Zeichen verwenden
			i++                    // Weiter zum nächsten Byte
		case r == 0x00: // Null-Byte ersetzen
			escaped.WriteRune('�') // oder ein anderes Ersatzzeichen
			i += size
		case r == 0xD6 && utf8.ValidString(string(r)): // 0xD6 0x79 als ungültige Sequenz behandeln
			escaped.WriteRune('�')
			i += size
		case r == '\'': // Einzelnes Hochkomma escapen
			escaped.WriteString("''")
			i += size
		case r == '\\': // Backslash escapen
			escaped.WriteString("\\\\")
			i += size
		default:
			escaped.WriteRune(r) // Andernfalls das Zeichen normal hinzufügen
			i += size
		}
	}

	return escaped.String()
}

// GetDomains retrieves the list of domains from the database.
func GetDomains(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`SELECT domain FROM public."DomainList"`)
	if err != nil {
		return nil, fmt.Errorf("error querying domains: %v", err)
	}
	defer rows.Close()

	var domains []string
	for rows.Next() {
		var domain string
		if err := rows.Scan(&domain); err != nil {
			return nil, fmt.Errorf("error scanning domain: %v", err)
		}
		domains = append(domains, domain)
	}

	return domains, nil
}

// GetDomains retrieves the list of domains from the database, filtered by even or odd rows.
func GetDomainsEvenUneven(db *sql.DB, even bool) ([]string, error) {
	var mod int
	if even {
		mod = 0 // Select even indexed rows
	} else {
		mod = 1 // Select odd indexed rows
	}

	query := `
        SELECT domain FROM (
            SELECT domain, ROW_NUMBER() OVER () as rownum
            FROM public."DomainList"
        ) AS subquery
        WHERE (rownum % 2) = $1;`

	rows, err := db.Query(query, mod)
	if err != nil {
		return nil, fmt.Errorf("error querying domains: %v", err)
	}
	defer rows.Close()

	var domains []string
	for rows.Next() {
		var domain string
		if err := rows.Scan(&domain); err != nil {
			return nil, fmt.Errorf("error scanning domain: %v", err)
		}
		domains = append(domains, domain)
	}

	return domains, nil
}

// GetDomainsByModulo retrieves the list of domains from the database, filtered by modulo and four categories
func GetDomainsByModulo(db *sql.DB, category int, modulo int) ([]string, error) {
	// Adjust the query to take modulo and category as parameters
	query := `
		SELECT domain
 		FROM public."DomainList" d
 		WHERE (d.id % $2) = $1;`

	// Run the query, passing category and modulo as parameters
	rows, err := db.Query(query, category, modulo)
	if err != nil {
		return nil, fmt.Errorf("error querying domains: %v", err)
	}
	defer rows.Close()

	var domains []string
	for rows.Next() {
		var domain string
		if err := rows.Scan(&domain); err != nil {
			return nil, fmt.Errorf("error scanning domain: %v", err)
		}
		domains = append(domains, domain)
	}

	return domains, nil
}

func GetUntestedDomainsByModulo(db *sql.DB, testCode string, category int, modulo int) ([]string, error) {
	// Adjust the query to exclude domains that have been tested with the given test_code
	query := `
        SELECT domain
        FROM public."DomainList" d
        WHERE NOT EXISTS (
                SELECT 1
                FROM public."DNSResults" r
                WHERE r.domain = d.domain
                AND r.test_code = $1
        )
        AND (d.id % $3) = $2;`

	// Run the query, passing testCode, category, and modulo as parameters
	rows, err := db.Query(query, testCode, category, modulo)
	if err != nil {
		return nil, fmt.Errorf("error querying domains: %v", err)
	}
	defer rows.Close()

	var domains []string
	for rows.Next() {
		var domain string
		if err := rows.Scan(&domain); err != nil {
			return nil, fmt.Errorf("error scanning domain: %v", err)
		}
		domains = append(domains, domain)
	}

	// Print the count of domains retrieved
	fmt.Printf("Retrieved %d domains\n", len(domains))

	return domains, nil
}

// GetDomainsForReassessmentByTestCode retrieves domains for reassessment based on the test code.
func GetDomainsForReassessmentByTestCode(db *sql.DB, testCode string) ([]string, error) {
	fmt.Println("Get Domains for Reassessment by TestCode")

	query := `
    SELECT domain
    FROM public."DomainList" d
    WHERE NOT EXISTS (
        SELECT 1
        FROM public."DNSResults" r
        WHERE r.domain = d.domain
        AND r.test_code = $1
    );`

	rows, err := db.Query(query, testCode)
	if err != nil {
		return nil, fmt.Errorf("error querying domains: %v", err)
	}
	defer rows.Close()

	var domains []string
	for rows.Next() {
		var domain string
		if err := rows.Scan(&domain); err != nil {
			return nil, fmt.Errorf("error scanning domain: %v", err)
		}
		domains = append(domains, domain)
	}

	fmt.Printf("Retrieved %d domains\n", len(domains))

	return domains, nil
}

// WriteResultsWithRetry writes the results to the database with retry logic.
func WriteResultsWithRetry(ctx context.Context, db *sql.DB, result DNSResult) error {
	var err error
	maxRetries := 8

	for i := 0; i < maxRetries; i++ {
		err = WriteDNSRecordResultsToDB(db, result)
		if err == nil {
			return nil // Success
		}

		if isTooManyConnectionsError(err) {
			// Exponential backoff: 2^i seconds
			backoffDuration := time.Duration(math.Pow(2, float64(i))) * time.Second

			// Use a timer with the context to handle cancellation
			timer := time.NewTimer(backoffDuration)
			defer timer.Stop() // Clean up the timer

			select {
			case <-ctx.Done():
				return ctx.Err() // Return the context error if cancelled
			case <-timer.C:
				fmt.Printf("Too many connections. Retrying in %v...\n", backoffDuration)
				continue
			}
		}

		// For other errors, do not retry
		break
	}

	return err
}

// isTooManyConnectionsError checks if the given error is a PostgreSQL "Too many connections" error.
func isTooManyConnectionsError(err error) bool {
	if pgErr, ok := err.(*pq.Error); ok {
		// Check for "Too many connections" error
		if pgErr.Code == "53300" { // PostgreSQL error code for "Too many connections"
			return true
		}
	}
	return false
}

func WriteDNSRecordResultsToDB(db *sql.DB, result DNSResult) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	//fmt.Printf("Writing Entries to Database, Batch Length: %v\n", result)// long debug
	fmt.Printf("Writing Entries to Database, Batch Length: %v\n", result.Domain)

	// Serialize the ErrorCode slice to JSON
	//fmt.Printf("Writing Domain to Database: %v\n", result) // Debug

	errorCodeJSON, err := json.Marshal(result.ErrorCode)
	HTTPSIPV4HINTJSON, err := json.Marshal(result.HTTPSIPV4HINT)
	HTTPSIPV6HINTJSON, err := json.Marshal(result.HTTPSIPV6HINT)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error marshalling ErrorCode: %v", err)
	}

	// Insert DNS Result
	var dnsResultID int64
	err = tx.QueryRow(`
			INSERT INTO public."DNSResults" 
			("domain", "https_raw", "https_ttl", "https_alpn", "https_ipv4_hint", "https_ipv6_hint", "https_ech_key", "querried_nameserver", "vantage_point", "error_msg", "test_code", "script_start_timestamp", "dns_query_start_timestamp", "dns_query_end_timestamp", "worker", "dial_timeout", "query_timeout", "max_attempts", "max_backoff_time", "test_date")
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
			RETURNING id`,
		result.Domain,
		result.HTTPSRAW,
		result.HTTPSTTL,
		result.HTTPSALPN,
		string(HTTPSIPV4HINTJSON),
		string(HTTPSIPV6HINTJSON),
		result.HTTPSECHKEY,
		string(result.QUERRIEDNS),
		result.VantagePoint,
		string(errorCodeJSON),
		result.TestCode,
		result.SCRIPTSTARTTIMESTAMP,
		result.DNSQUERYSTARTTIMESTAMP,
		result.DNSQUERYENDTIMESTAMP,
		result.WORKER,
		result.DIALTIMEOUT.Seconds(),
		result.QUERYTIMEOUT.Seconds(),
		result.MAXATTEMPTS,
		result.MAXBACKOFF.Seconds(),
		result.TESTDATE).Scan(&dnsResultID)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting into DNSResults: %v", err)
	}

	// Insert A records and link to DNSResults
	for i := range result.ARAW {
		aAsn := ""
		if len(result.AASN) > i {
			aAsn = result.AASN[i]
		}
		aAsnorg := ""
		if len(result.AASNORG) > i {
			aAsnorg = result.AASNORG[i]
		}
		aAsncity := ""
		if len(result.AASNCITY) > i {
			aAsncity = result.AASNCITY[i]
		}
		aAsncountry := ""
		if len(result.AASNCOUNTRY) > i {
			aAsncountry = result.AASNCOUNTRY[i]
		}

		var aRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."ARecords" 
				(raw, host, ttl, value, asn, asn_org, asn_city, asn_country, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
				RETURNING id`,
			result.ARAW[i],
			result.AHOST[i],
			result.ATTL[i],
			result.AVALUE[i],
			aAsn,
			aAsnorg,
			aAsncity,
			aAsncountry,
			result.ASIGNED,
			result.AVERIFIED,
			result.AVERIFICATIONERROR,
			result.AERRORMSG,
			result.ARTT,
			result.AATTEMPTS,
			result.ANETWORK,
			result.ARRSIGERRORMSG,
			result.AVALIDPERIOD,
			result.ARCODE,
			result.TESTDATE).Scan(&aRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting A record: %v", err)
		}

		_, err = tx.Exec(`
				INSERT INTO public."DNSResultsARecords" (dns_result_id, a_record_id)
				VALUES ($1, $2)`,
			dnsResultID, aRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking A record: %v", err)
		}
	}

	// Insert AAAA records and link to DNSResults
	for i := range result.AAAARAW {
		aaaaAsn := ""
		if len(result.AAAAASN) > i {
			aaaaAsn = result.AAAAASN[i]
		}
		aaaaAsnorg := ""
		if len(result.AAAAASNORG) > i {
			aaaaAsnorg = result.AAAAASNORG[i]
		}
		aaaaAsncity := ""
		if len(result.AAAAASNCITY) > i {
			aaaaAsncity = result.AAAAASNCITY[i]
		}
		aaaaAsncountry := ""
		if len(result.AAAAASNCOUNTRY) > i {
			aaaaAsncountry = result.AAAAASNCOUNTRY[i]
		}

		var aaaaRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."AAAARecords" 
				(raw, host, ttl, value, asn, asn_org, asn_city, asn_country, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
				RETURNING id`,
			result.AAAARAW[i],
			result.AAAAHOST[i],
			result.AAAATTL[i],
			result.AAAAVALUE[i],
			aaaaAsn,
			aaaaAsnorg,
			aaaaAsncity,
			aaaaAsncountry,
			result.AAAASIGNED,
			result.AAAAVERIFIED,
			result.AAAAVERIFICATIONERROR,
			result.AAAAERRORMSG,
			result.AAAARTT,
			result.AAAAATTEMPTS,
			result.AAAANETWORK,
			result.AAAARRSIGERRORMSG,
			result.AAAAVALIDPERIOD,
			result.AAAARCODE,
			result.TESTDATE).Scan(&aaaaRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting AAAA record: %v", err)
		}

		_, err = tx.Exec(`
				INSERT INTO public."DNSResultsAAAARecords" (dns_result_id, aaaa_record_id)
				VALUES ($1, $2)`,
			dnsResultID, aaaaRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking AAAA record: %v", err)
		}
	}

	// Insert NS records and link to DNSResults
	for i := range result.NSRAW {
		nsip := ""
		if len(result.NSIP) > i {
			nsip = result.NSIP[i]
		}
		nsipversion := ""
		if len(result.NSIPVERSION) > i {
			nsipversion = result.NSIPVERSION[i]
		}
		nsasn := ""
		if len(result.NSASN) > i {
			nsasn = result.NSASN[i]
		}
		nsasnorg := ""
		if len(result.NSASNORG) > i {
			nsasnorg = result.NSASNORG[i]
		}
		nsasncity := ""
		if len(result.NSASNCITY) > i {
			nsasncity = result.NSASNCITY[i]
		}
		nsasncountry := ""
		if len(result.NSASNCOUNTRY) > i {
			nsasncountry = result.NSASNCOUNTRY[i]
		}

		var nsRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."AuthoritativeNameserverRecords" 
				(raw, host, ttl, value, ip, ip_version, asn, asn_org, asn_city, asn_country, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
				RETURNING id`,
			result.NSRAW[i],
			result.NSHOST[i],
			result.NSTTL[i],
			result.NSVALUE[i],
			nsip,
			nsipversion,
			nsasn,
			nsasnorg,
			nsasncity,
			nsasncountry,
			result.NSSIGNED,
			result.NSVERIFIED,
			result.NSVERIFICATIONERROR,
			result.NSERRORMSG,
			result.NSRTT,
			result.NSATTEMPTS,
			result.NSNETWORK,
			result.NSRRSIGERRORMSG,
			result.NSVALIDPERIOD,
			result.NSRCODE,
			result.TESTDATE).Scan(&nsRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting NS record: %v", err)
		}

		_, err = tx.Exec(`
				INSERT INTO public."DNSResultsAuthoritativeNameserverRecords" (dns_result_id, authoritative_nameserver_record_id)
				VALUES ($1, $2)`,
			dnsResultID, nsRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking NS record: %v", err)
		}
	}

	// Insert TXT records and link to DNSResults
	for i := range result.TXTRAW {
		var txtRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."TXTRecords" 
				(raw, host, ttl, value, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
				RETURNING id`,
			result.TXTRAW[i],
			result.TXTHOST[i],
			result.TXTTTL[i],
			result.TXTVALUE[i],
			result.TXTSIGNED,
			result.TXTVERIFIED,
			result.TXTVERIFICATIONERROR,
			result.TXTERRORMSG,
			result.TXTRTT,
			result.TXTATTEMPTS,
			result.TXTNETWORK,
			result.TXTRRSIGERRORMSG,
			result.TXTVALIDPERIOD,
			result.TXTRCODE,
			result.TESTDATE).Scan(&txtRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting TXT record: %v", err)
		}

		_, err = tx.Exec(`
				INSERT INTO public."DNSResultsTXTRecords" (dns_result_id, txt_record_id)
				VALUES ($1, $2)`,
			dnsResultID, txtRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking TXT record: %v", err)
		}
	}
	// Insert MX records and link to DNSResults
	for i := range result.MXRAW {
		var mxRecordID int64

		mxip := ""
		if len(result.MXIP) > i {
			mxip = result.MXIP[i]
		}
		mxipversion := ""
		if len(result.MXIPVERSION) > i {
			mxipversion = result.MXIPVERSION[i]
		}
		mxasn := ""
		if len(result.MXASN) > i {
			mxasn = result.MXASN[i]
		}
		mxasnorg := ""
		if len(result.MXASNORG) > i {
			mxasnorg = result.MXASNORG[i]
		}
		mxasncity := ""
		if len(result.MXASNCITY) > i {
			mxasncity = result.MXASNCITY[i]
		}
		mxasncountry := ""
		if len(result.MXASNCOUNTRY) > i {
			mxasncountry = result.MXASNCOUNTRY[i]
		}

		err = tx.QueryRow(`
				INSERT INTO public."MXRecords" 
				(raw, host, ttl, value, ip, ip_version, asn, asn_org, asn_city, asn_country, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
				RETURNING id`,
			result.MXRAW[i],
			result.MXHOST[i],
			result.MXTTL[i],
			result.MXVALUE[i],
			mxip,
			mxipversion,
			mxasn,
			mxasnorg,
			mxasncity,
			mxasncountry,
			result.MXSIGNED,
			result.MXVERIFIED,
			result.MXVERIFICATIONERROR,
			result.MXERRORMSG,
			result.MXRTT,
			result.MXATTEMPTS,
			result.MXNETWORK,
			result.MXRRSIGERRORMSG,
			result.MXVALIDPERIOD,
			result.MXRCODE,
			result.TESTDATE).Scan(&mxRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting MX record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsMXRecords" (dns_result_id, mx_record_id) VALUES ($1, $2)`, dnsResultID, mxRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking MX record: %v", err)
		}
	}

	// Insert CNAME records and link to DNSResults
	for i := range result.CNAMERAW {
		var cnameRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."CNameRecords" 
				(raw, host, ttl, value, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
				RETURNING id`,
			result.CNAMERAW[i],
			result.CNAMEHOST[i],
			result.CNAMETTL[i],
			result.CNAMEVALUE[i],
			result.CNAMESIGNED,
			result.CNAMEVERIFIED,
			result.CNAMEVERIFICATIONERROR,
			result.CNAMEERRORMSG,
			result.CNAMERTT,
			result.CNAMEATTEMPTS,
			result.CNAMENETWORK,
			result.CNAMERRSIGERRORMSG,
			result.CNAMEVALIDPERIOD,
			result.CNAMERCODE,
			result.TESTDATE).Scan(&cnameRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting CNAME record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsCNameRecords" (dns_result_id, c_name_record_id) VALUES ($1, $2)`, dnsResultID, cnameRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking CNAME record: %v", err)
		}
	}

	// Insert RRSIG records and link to DNSResults
	for i := range result.RRSIGRAW {
		var rrsigRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."RRSIGRecords" 
				(raw, name, ttl, class, rr_type, rd_length, expiration, inception, key_tag, key_tag_string, type_covered, type_covered_string, labels, algorithm, algorithm_string, orig_ttl, signature, signer_name, time_valid_period, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30)
				RETURNING id`,
			result.RRSIGRAW[i],
			result.RRSIGNAME[i],
			result.RRSIGTTL[i],
			result.RRSIGCLASS[i],
			result.RRSIGRRTYPE[i],
			result.RRSIGRDLEN[i],
			result.RRSIGEXPIRATION[i],
			result.RRSIGINCEPTION[i],
			result.RRSIGKEYTAG[i],
			result.RRSIGKEYTAGSTRING[i],
			result.RRSIGTYPECOVERED[i],
			result.RRSIGTYPECOVEREDSTRING[i],
			result.RRSIGLABELS[i],
			result.RRSIGALGORITHM[i],
			result.RRSIGALGORITHMSTRING[i],
			result.RRSIGORIGTTL[i],
			result.RRSIGSIGNATURE[i],
			result.RRSIGSIGNERNAME[i],
			result.RRSIGTIMEVALIDPERIOD[i],
			result.RRSIGSIGNED,
			result.RRSIGVERIFIED,
			result.RRSIGVERIFICATIONERROR,
			result.RRSIGERRORMSG,
			result.RRSIGRTT,
			result.RRSIGATTEMPTS,
			result.RRSIGNETWORK,
			result.RRSIGRRSIGERRORMSG,
			result.RRSIGVALIDPERIOD,
			result.RRSIGRCODE,
			result.TESTDATE).Scan(&rrsigRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting RRSIG record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsRRSIGRecords" (dns_result_id, rrsig_record_id) VALUES ($1, $2)`, dnsResultID, rrsigRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking RRSIG record: %v", err)
		}
	}

	// Insert DNSKEY records and link to DNSResults
	for i := range result.DNSKEYRAW {
		var dnskeyRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."DNSKeyRecords"
				(raw, name, ttl, class, rr_type, rd_length, algorithm, algorithm_string, key_tag, key_tag_string, flags, public_key, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
				RETURNING id`,
			result.DNSKEYRAW[i],
			result.DNSKEYNAME[i],
			result.DNSKEYTTL[i],
			result.DNSKEYCLASS[i],
			result.DNSKEYRRTYPE[i],
			result.DNSKEYRDLEN[i],
			result.DNSKEYALGORITHM[i],
			result.DNSKEYALGORITHMSTRING[i],
			result.DNSKEYKEYTAG[i],
			result.DNSKEYKEYTAGSTRING[i],
			result.DNSKEYFLAGS[i],
			result.DNSKEYPUBLICKEY[i],
			result.DNSKEYSIGNED,
			result.DNSKEYVERIFIED,
			result.DNSKEYVERIFICATIONERROR,
			result.DNSKEYERRORMSG,
			result.DNSKEYRTT,
			result.DNSKEYATTEMPTS,
			result.DNSKEYNETWORK,
			result.DNSKEYRRSIGERRORMSG,
			result.DNSKEYVALIDPERIOD,
			result.DNSKEYRCODE,
			result.TESTDATE).Scan(&dnskeyRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting DNSKEY record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsDNSKeyRecords" (dns_result_id, dns_key_record_id) VALUES ($1, $2)`, dnsResultID, dnskeyRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking DNSKEY record: %v", err)
		}
	}

	// Insert CAA records and link to DNSResults
	for i := range result.CAARAW {
		var caaRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."CAARecords"
				(raw, name, ttl, class, rr_type, rd_length, flag, value, tag, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
				RETURNING id`,
			EscapeSpecialChars(result.CAARAW[i]),
			EscapeSpecialChars(result.CAANAME[i]),
			result.CAATTL[i],
			result.CAACLASS[i],
			result.CAARRTYPE[i],
			result.CAARDLEN[i],
			result.CAAFLAG[i],
			EscapeSpecialChars(result.CAAVALUE[i]),
			result.CAATAG[i],
			result.CAASIGNED,
			result.CAAVERIFIED,
			EscapeSpecialChars(result.CAAVERIFICATIONERROR),
			EscapeSpecialChars(result.CAAERRORMSG),
			result.CAARTT,
			result.CAAATTEMPTS,
			result.CAANETWORK,
			result.CAARRSIGERRORMSG,
			result.CAAVALIDPERIOD,
			result.CAARCODE,
			result.TESTDATE).Scan(&caaRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting CAA record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsCAARecords" (dns_result_id, caa_record_id) VALUES ($1, $2)`, dnsResultID, caaRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking CAA record: %v", err)
		}
	}

	// Insert SOA records and link to DNSResults
	for i := range result.SOARAW {
		var soaRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."SOARecords"
				(raw, name, ttl, class, rr_type, rd_length, ns, mbox, refresh, refresh_string, retry, serial, min_ttl, expire, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25)
				RETURNING id`,
			result.SOARAW[i],
			result.SOANAME[i],
			result.SOATTL[i],
			result.SOACLASS[i],
			result.SOARRTYPE[i],
			result.SOARDLEN[i],
			result.SOANS[i],
			result.SOAMBOX[i],
			result.SOAREFRESH[i],
			result.SOAREFRESHSTRING[i],
			result.SOARETRY[i],
			result.SOASERIAL[i],
			result.SOAMINTTL[i],
			result.SOAEXPIRE[i],
			result.SOASIGNED,
			result.SOAVERIFIED,
			result.SOAVERIFICATIONERROR,
			result.SOAERRORMSG,
			result.SOARTT,
			result.SOAATTEMPTS,
			result.SOANETWORK,
			result.SOARRSIGERRORMSG,
			result.SOAVALIDPERIOD,
			result.SOARCODE,
			result.TESTDATE).Scan(&soaRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting SOA record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsSOARecords" (dns_result_id, soa_record_id) VALUES ($1, $2)`, dnsResultID, soaRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking SOA record: %v", err)
		}
	}

	// Insert DS records and link to DNSResults
	for i := range result.DSRAW {
		var dsRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."DSRecords" 
				(raw, name, ttl, class, rr_type, rd_length, algorithm, algorithm_string, key_tag, key_tag_string, digest, digest_type, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
				RETURNING id`,
			result.DSRAW[i],
			result.DSNAME[i],
			result.DSTTL[i],
			result.DSCLASS[i],
			result.DSRRTYPE[i],
			result.DSRDLEN[i],
			result.DSALGORITHM[i],
			result.DSALGORITHMSTRING[i],
			result.DSKEYTAG[i],
			result.DSKEYTAGSTRING[i],
			result.DSDIGEST[i],
			result.DSDIGESTTYPE[i],
			result.DSSIGNED,
			result.DSVERIFIED,
			result.DSVERIFICATIONERROR,
			result.DSERRORMSG,
			result.DSRTT,
			result.DSATTEMPTS,
			result.DSNETWORK,
			result.DSRRSIGERRORMSG,
			result.DSVALIDPERIOD,
			result.DSRCODE,
			result.TESTDATE).Scan(&dsRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting DS record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsDSRecords" (dns_result_id, ds_record_id) VALUES ($1, $2)`, dnsResultID, dsRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking DS record: %v", err)
		}
	}

	// Insert CDS records and link to DNSResults
	for i := range result.CDSRAW {
		var cdsRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."CDSRecords" 
				(raw, name, ttl, class, rr_type, rd_length, algorithm, algorithm_string, key_tag, key_tag_string, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
				RETURNING id`,
			result.CDSRAW[i],
			result.CDSNAME[i],
			result.CDSTTL[i],
			result.CDSCLASS[i],
			result.CDSRRTYPE[i],
			result.CDSRDLEN[i],
			result.CDSALGORITHM[i],
			result.CDSALGORITHMSTRING[i],
			result.CDSKEYTAG[i],
			result.CDSKEYTAGSTRING[i],
			result.CDSSIGNED,
			result.CDSVERIFIED,
			result.CDSVERIFICATIONERROR,
			result.CDSERRORMSG,
			result.CDSRTT,
			result.CDSATTEMPTS,
			result.CDSNETWORK,
			result.CDSRRSIGERRORMSG,
			result.CDSVALIDPERIOD,
			result.CDSRCODE,
			result.TESTDATE).Scan(&cdsRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting CDS record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsCDSRecords" (dns_result_id, cds_record_id) VALUES ($1, $2)`, dnsResultID, cdsRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking CDS record: %v", err)
		}
	}

	// Insert NSEC records and link to DNSResults
	for i := range result.NSECRAW {
		var nsecRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."NSECRecords"
				(raw, name, ttl, class, rr_type, rd_length, type_bit_map_string, next_domain, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
				RETURNING id`,
			result.NSECRAW[i],
			result.NSECNAME[i],
			result.NSECTTL[i],
			result.NSECCLASS[i],
			result.NSECRRTYPE[i],
			result.NSECRDLEN[i],
			result.NSECTYPEBITMAPSTRING[i],
			result.NSECNEXTDOMAIN[i],
			result.NSECSIGNED,
			result.NSECVERIFIED,
			result.NSECVERIFICATIONERROR,
			result.NSECERRORMSG,
			result.NSECRTT,
			result.NSECATTEMPTS,
			result.NSECNETWORK,
			result.NSECRRSIGERRORMSG,
			result.NSECVALIDPERIOD,
			result.NSECRCODE,
			result.TESTDATE).Scan(&nsecRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting NSEC record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsNSECRecords" (dns_result_id, nsec_record_id) VALUES ($1, $2)`, dnsResultID, nsecRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking NSEC record: %v", err)
		}
	}

	// Insert NSEC3 records and link to DNSResults
	for i := range result.NSEC3RAW {
		var nsec3RecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."NSEC3Records" 
				(raw, name, ttl, class, rr_type, rd_length, hash_length, next_domain, hash, iterations, salt, salt_length, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
				RETURNING id`,
			result.NSEC3RAW[i],
			result.NSEC3NAME[i],
			result.NSEC3TTL[i],
			result.NSEC3CLASS[i],
			result.NSEC3RRTYPE[i],
			result.NSEC3RDLEN[i],
			result.NSEC3HASHLENGTH[i],
			result.NSEC3NEXTDOMAIN[i],
			result.NSEC3HASH[i],
			result.NSEC3ITERATIONS[i],
			result.NSEC3SALT[i],
			result.NSEC3SALTLENGTH[i],
			result.NSEC3SIGNED,
			result.NSEC3VERIFIED,
			result.NSEC3VERIFICATIONERROR,
			result.NSEC3ERRORMSG,
			result.NSEC3RTT,
			result.NSEC3ATTEMPTS,
			result.NSEC3NETWORK,
			result.NSEC3RRSIGERRORMSG,
			result.NSEC3VALIDPERIOD,
			result.NSEC3RCODE,
			result.TESTDATE).Scan(&nsec3RecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting NSEC3 record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsNSEC3Records" (dns_result_id, nsec3_record_id) VALUES ($1, $2)`, dnsResultID, nsec3RecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking NSEC3 record: %v", err)
		}
	}

	// Insert NSEC3PARAM records and link to DNSResults
	for i := range result.NSEC3PARAMRAW {
		var nsec3paramRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."NSEC3PARAMRecords" 
				(raw, name, ttl, class, rr_type, rd_length, hash, iterations, salt, salt_length, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
				RETURNING id`,
			result.NSEC3PARAMRAW[i],
			result.NSEC3PARAMNAME[i],
			result.NSEC3PARAMTTL[i],
			result.NSEC3PARAMCLASS[i],
			result.NSEC3PARAMRRTYPE[i],
			result.NSEC3PARAMRDLEN[i],
			result.NSEC3PARAMHASH[i],
			result.NSEC3PARAMITERATIONS[i],
			result.NSEC3PARAMSALT[i],
			result.NSEC3PARAMSALTLENGTH[i],
			result.NSEC3PARAMSIGNED,
			result.NSEC3PARAMVERIFIED,
			result.NSEC3PARAMVERIFICATIONERROR,
			result.NSEC3PARAMERRORMSG,
			result.NSEC3PARAMRTT,
			result.NSEC3PARAMATTEMPTS,
			result.NSEC3PARAMNETWORK,
			result.NSEC3PARAMRRSIGERRORMSG,
			result.NSEC3PARAMVALIDPERIOD,
			result.NSEC3PARAMRCODE,
			result.TESTDATE).Scan(&nsec3paramRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting NSEC3PARAM record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsNSEC3PARAMRecords" (dns_result_id, nsec3param_record_id) VALUES ($1, $2)`, dnsResultID, nsec3paramRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking NSEC3PARAM record: %v", err)
		}
	}

	// Insert CDNSKEY records and link to DNSResults
	for i := range result.CDNSKEYRAW {
		var cdnskeyRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."CDNSKeyRecords"
				(raw, name, ttl, class, rr_type, rd_length, flags, key_tag, key_tag_string, protocol, dns_key, algorithm, algorithm_string, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24)
				RETURNING id`,
			result.CDNSKEYRAW[i],
			result.CDNSKEYNAME[i],
			result.CDNSKEYTTL[i],
			result.CDNSKEYCLASS[i],
			result.CDNSKEYRRTYPE[i],
			result.CDNSKEYRDLEN[i],
			result.CDNSKEYFLAGS[i],
			result.CDNSKEYKEYTAG[i],
			result.CDNSKEYKEYTAGSTRING[i],
			result.CDNSKEYPROTOCOL[i],
			result.CDNSKEYDNSKEY[i],
			result.CDNSKEYALGORITHM[i],
			result.CDNSKEYALGORITHMSTRING[i],
			result.CDNSKEYSIGNED,
			result.CDNSKEYVERIFIED,
			result.CDNSKEYVERIFICATIONERROR,
			result.CDNSKEYERRORMSG,
			result.CDNSKEYRTT,
			result.CDNSKEYATTEMPTS,
			result.CDNSKEYNETWORK,
			result.CDNSKEYRRSIGERRORMSG,
			result.CDNSKEYVALIDPERIOD,
			result.CDNSKEYRCODE,
			result.TESTDATE).Scan(&cdnskeyRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting CDNSKEY record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsCDNSKeyRecords" (dns_result_id, cdnskey_record_id) VALUES ($1, $2)`, dnsResultID, cdnskeyRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking CDNSKEY record: %v", err)
		}
	}

	// Insert TLSA records and link to DNSResults
	for i := range result.TLSARAW {
		var tlsaRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."TLSARecords"
				(raw, name, ttl, class, rr_type, rd_length, usage, certificate, matching_type, selector, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
				RETURNING id`,
			result.TLSARAW[i],
			result.TLSANAME[i],
			result.TLSATTL[i],
			result.TLSACLASS[i],
			result.TLSARRTYPE[i],
			result.TLSARDLEN[i],
			result.TLSAUSAGE[i],
			result.TLSACERTIFICATE[i],
			result.TLSAMATCHINGTYPE[i],
			result.TLSASELECTOR[i],
			result.TLSASIGNED,
			result.TLSAVERIFIED,
			result.TLSAVERIFICATIONERROR,
			result.TLSAERRORMSG,
			result.TLSARTT,
			result.TLSAATTEMPTS,
			result.TLSANETWORK,
			result.TLSARRSIGERRORMSG,
			result.TLSAVALIDPERIOD,
			result.TLSARCODE,
			result.TESTDATE).Scan(&tlsaRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting TLSA record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsTLSARecords" (dns_result_id, tlsa_record_id) VALUES ($1, $2)`, dnsResultID, tlsaRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking TLSA record: %v", err)
		}
	}

	// Insert OPT records and link to DNSResults
	for i := range result.OPTRAW {
		var optRecordID int64
		err = tx.QueryRow(`
				INSERT INTO public."OPTRecords"
				(raw, name, ttl, class, rr_type, rd_length, option, do, extended_r_code, version, udp_size, z, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23)
				RETURNING id`,
			result.OPTRAW[i],
			result.OPTNAME[i],
			result.OPTTTL[i],
			result.OPTCLASS[i],
			result.OPTRRTYPE[i],
			result.OPTRDLEN[i],
			result.OPTOPTION[i],
			result.OPTDO[i],
			result.OPTEXTENDEDRCODE[i],
			result.OPTVERSION[i],
			result.OPTUDPSIZE[i],
			result.OPTZ[i],
			result.OPTSIGNED,
			result.OPTVERIFIED,
			result.OPTVERIFICATIONERROR,
			result.OPTERRORMSG,
			result.OPTRTT,
			result.OPTATTEMPTS,
			result.OPTNETWORK,
			result.OPTRRSIGERRORMSG,
			result.OPTVALIDPERIOD,
			result.OPTRCODE,
			result.TESTDATE).Scan(&optRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting OPT record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsOPTRecords" (dns_result_id, opt_record_id) VALUES ($1, $2)`, dnsResultID, optRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking OPT record: %v", err)
		}
	}

	// Insert HTTPS records and link to DNSResults
	//for i := range result.HTTPSRAW {
	var httpsRecordID int64
	if result.HTTPSRAW != "" {
		err = tx.QueryRow(`
				INSERT INTO public."HTTPSRecords"
				(domain, raw, name, ttl, class, rr_type, rd_length, priority, target, mandatory, alpn, ipv4_hint, ipv6_hint, ech_config, doh_target, "0_doh_config", esni_keys, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28)
				RETURNING id`,
			result.Domain,
			result.HTTPSRAW,
			result.HTTPSNAME,
			result.HTTPSTTL,
			result.HTTPSCLASS,
			result.HTTPSRRTYPE,
			result.HTTPSRDLEN,
			result.HTTPSPRIORITY,
			result.HTTPSTARGET,
			result.HTTPSMANDATORY,
			result.HTTPSALPN,
			result.HTTPSIPV4HINT,
			result.HTTPSIPV6HINT,
			result.HTTPSECHKEY,
			result.HTTPSDOHTTARGET,
			result.HTTPSODOHCONFIG,
			result.HTTPSESNIKEYS,
			result.HTTPSSIGNED,
			result.HTTPSVERIFIED,
			result.HTTPSVERIFICATIONERROR,
			result.HTTPSERRORMSG,
			result.HTTPSRTT,
			result.HTTPSATTEMPTS,
			result.HTTPSNETWORK,
			result.HTTPSRRSIGERRORMSG,
			result.HTTPSVALIDPERIOD,
			result.HTTPSRCODE,
			result.TESTDATE).Scan(&httpsRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting HTTPS record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsHTTPSRecords" (dns_result_id, https_record_id) VALUES ($1, $2)`, dnsResultID, httpsRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking HTTPS record: %v", err)
		}

		// Insert ECH records and link to DNSResults, HTTPSRecords and ECH sub records
		for i := range result.ECHCONFIGB64 {
			fmt.Printf("Inserting ECH record for domain %s\n", result.Domain)
			fmt.Printf("ECHCONFIGB64: %s\n", result.ECHCONFIGB64[i])
			var echConfigID int64

			err = tx.QueryRow(`
	INSERT INTO public."ECHConfigs"
	(domain, test_code, test_date, ech_config_base64, ech_config_raw, ech_config_version, ech_config_length, ech_config_id, ech_kem_id, ech_public_key, ech_max_name_length, ech_public_name)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	RETURNING id`,
				result.Domain,
				result.TestCode,
				result.TESTDATE,
				result.ECHCONFIGB64[i],
				result.ECHCONFIGRAW[i],
				result.ECHCONFIGTLSVERSION[i],
				result.ECHCONFIGLENGTH[i],
				result.ECHCONFIGID[i],
				result.ECHKEMID[i],
				result.ECHPUBLICKEY[i],
				result.ECHMAXNAMELENGTH[i],
				result.ECHPUBLICNAME[i]).Scan(&echConfigID)

			if err != nil {
				tx.Rollback()
				return fmt.Errorf("error inserting ECHConfigs: %v", err)
			}

			_, err = tx.Exec(`INSERT INTO public."DNSResultsECHConfigs" (dns_result_id, ech_config_id) VALUES ($1, $2)`, dnsResultID, echConfigID)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("error linking ECHConfigs to DNSResults: %v", err)
			}
			//Linking to ECHConfigs to HTTPSRecords as one HTTPSRecord can contain multiple SVCB Records
			_, err = tx.Exec(`INSERT INTO public."HTTPSRecordsECHConfigs" (https_record_id, ech_config_id) VALUES ($1, $2)`, httpsRecordID, echConfigID)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("error linking ECHConfigs to HTTPSRecords: %v", err)
			}

			for j := range result.ECHCIPHERKDFID {
				var echCipherID int64

				err = tx.QueryRow(`
				INSERT INTO public."ECHCiphers"
				(domain, test_code, test_date, ech_config_base64, ech_cipher_kdf_id, ech_cipher_aead_id)
				VALUES ($1, $2, $3, $4, $5, $6)
				RETURNING id`,
					result.Domain,
					result.TestCode,
					result.TESTDATE,
					result.ECHCONFIGB64[i],
					result.ECHCIPHERKDFID[j],
					result.ECHCIPHERAEADID[j]).Scan(&echCipherID)

				if err != nil {
					tx.Rollback()
					return fmt.Errorf("error inserting ECHCiphers: %v", err)
				}

				_, err = tx.Exec(`INSERT INTO public."ECHConfigsECHCiphers" (ech_config_id, ech_cipher_id) VALUES ($1, $2)`, echConfigID, echCipherID)
				if err != nil {
					tx.Rollback()
					return fmt.Errorf("error linking ECHCiphers to ECHConfigs: %v", err)
				}
			}
			for k := range result.ECHEXTENSIONTYPE {
				var echExtensionID int64
				err = tx.QueryRow(`
				INSERT INTO public."ECHExtensions"
				(domain, test_code, test_date, ech_config_base64, ech_extension_type, ech_extension_data)
				VALUES ($1, $2, $3, $4, $5, $6)
				RETURNING id`,
					result.Domain,
					result.TestCode,
					result.TESTDATE,
					result.ECHCONFIGB64[i],
					result.ECHEXTENSIONTYPE[k],
					result.ECHEXTENSIONDATA[k]).Scan(&echExtensionID)

				if err != nil {
					tx.Rollback()
					return fmt.Errorf("error inserting ECHExtensions: %v", err)
				}

				_, err = tx.Exec(`INSERT INTO public."ECHConfigsECHExtensions" (ech_config_id, ech_extension_id) VALUES ($1, $2)`, echConfigID, echExtensionID)
				if err != nil {
					tx.Rollback()
					return fmt.Errorf("error linking ECHExtensions to ECHConfigs: %v", err)
				}
			}
		}
	}

	// Insert SVCB records and link to DNSResults
	for i := range result.SVCBRAW {
		var svcbRecordID int64
		err = tx.QueryRow(`
					INSERT INTO public."SVCBRecords"
					(domain, raw, name, target, ttl, class, rr_type, rd_length, priority, value, signed, verified, verification_error, error_msg, "rtt_ns", attempt, network, "rrsig_error_msg", "rrsig_valid_period", "r_code", "test_date")
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
					RETURNING id`,
			result.Domain,
			result.SVCBRAW[i],
			result.SVCBNAME[i],
			result.SVCBTARGET[i],
			result.SVCBTTL[i],
			result.SVCBCLASS[i],
			result.SVCBRRTYPE[i],
			result.SVCBRDLEN[i],
			result.SVCBPRIORITY[i],
			result.SVCBVALUE[i],
			result.SVCBSIGNED,
			result.SVCBVERIFIED,
			result.SVCBVERIFICATIONERROR,
			result.SVCBERRORMSG,
			result.SVCBRTT,
			result.SVCBATTEMPTS,
			result.SVCBNETWORK,
			result.SVCBRRSIGERRORMSG,
			result.SVCBVALIDPERIOD,
			result.SVCBRCODE,
			result.TESTDATE).Scan(&svcbRecordID)

		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting SVCB record: %v", err)
		}

		_, err = tx.Exec(`INSERT INTO public."DNSResultsSVCBRecords" (dns_result_id, svcb_record_id) VALUES ($1, $2)`, dnsResultID, svcbRecordID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error linking SVCB record: %v", err)
		}
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}
