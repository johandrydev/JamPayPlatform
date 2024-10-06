-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS audit_trail
(
    id
    SERIAL
    PRIMARY
    KEY,
    event_type
    VARCHAR
(
    10
) NOT NULL,
    event_description TEXT NOT NULL,
    table_name VARCHAR
(
    255
) NOT NULL,
    old_data JSONB,
    new_data JSONB,
    user_id INTEGER,
    timestamp TIMESTAMP NOT NULL
    );

CREATE
OR REPLACE FUNCTION log_audit_trail()
RETURNS TRIGGER AS $$
BEGIN
    -- Insert the audit log entry
INSERT INTO audit_trail (event_type,
                         event_description,
                         table_name,
                         old_data,
                         new_data,
                         user_id,
                         timestamp)
VALUES (TG_OP, -- TG_OP is the type of operation: INSERT, UPDATE, DELETE
        'Change in ' || TG_TABLE_NAME || ' by user ' || COALESCE(current_user, 'unknown'),
        TG_TABLE_NAME,
        row_to_json(OLD), -- OLD data (before update/delete) in JSON format
        row_to_json(NEW), -- NEW data (after insert/update) in JSON format
        NULL,
        NOW());

RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER payment_audit_trigger
    AFTER INSERT OR
UPDATE OR
DELETE
ON payments
    FOR EACH ROW EXECUTE FUNCTION log_audit_trail();

CREATE TRIGGER merchant_audit_trigger
    AFTER INSERT OR
UPDATE OR
DELETE
ON merchants
    FOR EACH ROW EXECUTE FUNCTION log_audit_trail();

CREATE TRIGGER customer_audit_trigger
    AFTER INSERT OR
UPDATE OR
DELETE
ON customers
    FOR EACH ROW EXECUTE FUNCTION log_audit_trail();

---- create above / drop below ----

DROP FUNCTION IF EXISTS log_audit_trail() CASCADE;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.