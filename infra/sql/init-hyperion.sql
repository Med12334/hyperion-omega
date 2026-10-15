CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE IF NOT EXISTS tenant_compute_quotas (
    tenant_id VARCHAR(64) PRIMARY KEY,
    organization_name VARCHAR(120) NOT NULL,
    allocated_vcpu INTEGER NOT NULL,
    allocated_gpu_units INTEGER NOT NULL,
    consumed_vcpu_seconds BIGINT NOT NULL DEFAULT 0,
    credit_balance_usd NUMERIC(14, 4) NOT NULL DEFAULT 1000.0000,
    version BIGINT NOT NULL DEFAULT 0,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS merkle_governance_ledger (
    audit_id UUID PRIMARY KEY,
    tenant_id VARCHAR(64) NOT NULL,
    cluster_region VARCHAR(32) NOT NULL,
    operation_type VARCHAR(64) NOT NULL,
    previous_hash VARCHAR(64) NOT NULL,
    merkle_hash VARCHAR(64) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS outbox_events (
    event_id UUID PRIMARY KEY,
    aggregate_type VARCHAR(50) NOT NULL,
    aggregate_id VARCHAR(64) NOT NULL,
    event_type VARCHAR(50) NOT NULL,
    payload TEXT NOT NULL,
    processed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
