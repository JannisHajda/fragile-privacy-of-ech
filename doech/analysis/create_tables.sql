CREATE TABLE IF NOT EXISTS dns_results
(
    worker_id String,
    run_uuid String,
    domain String,
    start DateTime64(3, 'UTC'),
    end DateTime64(3, 'UTC'),
    dns_a_results String,
    dns_aaaa_results String,
    dns_svcb_results String,
    dns_https_results String
)
ENGINE = MergeTree()
ORDER BY (worker_id, run_uuid, start);

CREATE TABLE IF NOT EXISTS doech_results
(
    worker_id String,
    run_uuid String,
    domain String,
    start DateTime64(3, 'UTC'),
    end DateTime64(3, 'UTC'),
    results String
)
ENGINE = MergeTree()
ORDER BY (worker_id, run_uuid, start);

