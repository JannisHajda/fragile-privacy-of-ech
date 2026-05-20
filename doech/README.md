# doech
Developer: Jannis Hajda

This directory contains the code for reproducing the active browser measurements using our custom browser extension.

### Running Measurements
The source code for reproducing the measurements is available inside the `src` directory. 
The measurements are divided into two steps:
1. Determining the number of domains that advertise ECH support via DNS (`dns_test.py`)
2. Performing active analysis of DoH and ECH usage on a per-request level to validate the privacy promises of ECH (`doech_test.py`)

Make sure to start the ClickHouse database (see `docker-compose.yml`) and ensure that you have Geckodriver installed locally. Don't forget to adjust the configurations in the specific files to your needs.

### Performing Analysis
The analysis code is provided within the `analysis` directory. 
Again, it is separated into `analysis_dns` and `analysis_doech`. 
To enable reproducibility, we have included aggregated datasets for the `analysis_doech` part in the `results` directory.
