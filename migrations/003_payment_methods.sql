-- Write your migrate up statements here

CREATE TABLE payment_methods
(
    id              uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    external_id     VARCHAR(255),
    owner_id        uuid         NOT NULL,
    type            VARCHAR(255) NOT NULL,
    product_number  VARCHAR(255) NOT NULL,
    expiration_date DATE,
    created_at      TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE payment_methods
    ADD CONSTRAINT fk_payment_methods_owner_id FOREIGN KEY (owner_id) REFERENCES customers (id);

BEGIN;

INSERT INTO payment_methods (owner_id, type, product_number, expiration_date, created_at)
VALUES
    ((SELECT id FROM customers WHERE email = 'pablo@client.com'), 'CREDIT_CARD', '4242424242424242', '2025-12-31', now()),
    ((SELECT id FROM customers WHERE email = 'maria@example.com'),'DEBIT_CARD', '4242424242424242', '2024-11-30', now()),

COMMIT;

---- create above / drop below ----

DROP TABLE payment_methods;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
