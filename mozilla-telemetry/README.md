# Mozilla Telemetry
Developer: John Bauer

This directory contains the code and resources for reproducing the client-side adoption analysis using aggregated Mozilla Firefox telemetry.

### Fetching the Data
The raw telemetry data used in this analysis is owned by Mozilla. 
The dataset is publicly accessible via Mozilla's BigQuery project. 
To fetch the data, run the following Standard SQL query in the Google Cloud Console:

```sql
SELECT 
    submission_date, metric, label, key, country_code, 
    handshakes, total_client_count
FROM `mozilla-public-data.telemetry_derived.ech_adoption_rate_v1`
WHERE submission_date BETWEEN DATE '2025-02-18' AND DATE '2025-09-14';
```

### Analysis
Our analysis code is available inside the `analysis.ipynb` notebook.
