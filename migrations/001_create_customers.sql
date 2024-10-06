-- Write your migrate up statements here
CREATE TABLE users
(
    id              uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    email           VARCHAR(255) NOT NULL UNIQUE,
    role            VARCHAR(255) NOT NULL,
    hashed_password VARCHAR(255) NOT NULL,
    created_at      TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP        DEFAULT NULL
);

CREATE TABLE customers
(
    id         uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    external_id VARCHAR(255),
    name       VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL UNIQUE,
    status     VARCHAR(255) NOT NULL,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE customers
    ADD CONSTRAINT fk_users_email FOREIGN KEY (email) REFERENCES users (email);

INSERT INTO users (email, role, hashed_password, created_at)
VALUES ('pablo@client.com', 'CUSTOMER', 'RANDOMHASH', now()),
       ('maria@example.com', 'CUSTOMER', 'HASH1', now()),
       ('pedro@example.com', 'CUSTOMER', 'HASH2', now());

INSERT INTO customers (name, email, status, created_at)
VALUES ('Pablo Mendez', 'pablo@client.com', 'ACTIVE', now()),
       ('Maria Perez', 'maria@example.com', 'ACTIVE', now()),
       ('Pedro Rodriguez', 'pedro@example.com', 'INACTIVE', now());


---- create above / drop below ----

DROP TABLE customers;
DROP TABLE users;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
