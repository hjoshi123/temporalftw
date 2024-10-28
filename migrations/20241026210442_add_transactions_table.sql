CREATE TABLE transaction_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO transaction_types (name) VALUES ('credit');
INSERT INTO transaction_types (name) VALUES ('debit');

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    amount DECIMAL(10, 2) NOT NULL,
    description TEXT NOT NULL,
    transaction_type_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    from_account_number BIGINT NOT NULL,
    to_account_number BIGINT NOT NULL,
    FOREIGN KEY (from_account_number) REFERENCES accounts (account_number),
    FOREIGN KEY (to_account_number) REFERENCES accounts (account_number),
    FOREIGN KEY (transaction_type_id) REFERENCES transaction_types (id),
    CHECK (amount > 0)
);