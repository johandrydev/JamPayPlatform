-- Write your migrate up statements here
CREATE TABLE merchants
(
    id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    email        VARCHAR(255) NOT NULL UNIQUE,
    name         VARCHAR(255) NOT NULL,
    bank_account VARCHAR(255) NOT NULL,
    status       VARCHAR(255) NOT NULL,
    created_at   timestamptz(6) NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    updated_at   timestamptz(6) NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
ALTER TABLE merchants
    ADD CONSTRAINT fk_users_email FOREIGN KEY (email) REFERENCES users (email);

INSERT INTO users (email, role, hashed_password, created_at)
VALUES ('peach@mail.com', 'MERCHANT', '$2a$10$WZvWSnML1cnDm.VSZhWjGuQM9eQIVQRdMr5cD2sFHhF9yhB27cFxi', now()),
         ('mario@mail.com', 'MERCHANT', '$2a$10$WZvWSnML1cnDm.VSZhWjGuQM9eQIVQRdMr5cD2sFHhF9yhB27cFxi', now()),
         ('luigi@mail.com', 'MERCHANT', '$2a$10$WZvWSnML1cnDm.VSZhWjGuQM9eQIVQRdMr5cD2sFHhF9yhB27cFxi', now());

INSERT INTO merchants (email, name, bank_account, status, created_at)
VALUES ('peach@mail.com', 'Peach Bros', '123456789', 'VERIFIED', now()),
       ('mario@mail.com', 'Mario Bros', '987654321', 'VERIFIED', now()),
       ('luigi@mail.com', 'Luigi Bros', '123123123', 'VERIFIED', now());

---- create above / drop below ----

DROP TABLE merchants;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
