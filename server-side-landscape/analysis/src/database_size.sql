SELECT
    table_name,
    ROUND(SUM(data_length + index_length) / 1024 / 1024, 2) AS table_size_mb
FROM
    information_schema.tables
WHERE
    table_schema = 'thesis_ech'
    AND table_name IN ('TxtRecords', 'MxRecords', 'ARecords', 'AAAARecords', 'CNameRecords', 'DomainResults', 'AuthoritativeNameservers', 'DNSResultsAAAARecords', 'DNSResultsARecords', 'DNSResultsAuthoritativeNameservers', 'DNSResultsCNameRecords', 'DNSResultsMxRecords', 'DNSResultsTxtRecords' )  -- specify your table names
GROUP BY
    table_name;

SELECT
    ROUND(SUM(data_length + index_length) / 1024 / 1024, 2) AS total_size_mb
FROM
    information_schema.tables
WHERE
    table_schema = 'thesis_ech'
    AND table_name IN ('TxtRecords', 'MxRecords', 'ARecords', 'AAAARecords', 'CNameRecords', 'DomainResults', 'AuthoritativeNameservers', 'DNSResultsAAAARecords', 'DNSResultsARecords', 'DNSResultsAuthoritativeNameservers', 'DNSResultsCNameRecords', 'DNSResultsMxRecords', 'DNSResultsTxtRecords' );  -- specify your table names
