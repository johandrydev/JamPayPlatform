-- Write your migrate up statements here

CREATE TABLE payments
(
    id                uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    external_id       VARCHAR(255),
    merchant_id       uuid             NOT NULL,
    customer_id       uuid             NOT NULL,
    payment_method_id uuid             NOT NULL,
    amount            double precision NOT NULL CHECK (amount >= 0),
    status            VARCHAR(255)     NOT NULL,
    created_at        TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    processed_at      TIMESTAMP
);

---- create above / drop below ----

DROP TABLE payments;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
