CREATE TABLE transaction_statuses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO transaction_statuses (name) VALUES ('pending');
INSERT INTO transaction_statuses (name) VALUES ('completed');
INSERT INTO transaction_statuses (name) VALUES ('failed');

ALTER TABLE transactions ADD COLUMN transaction_status_id INT NOT NULL DEFAULT 1;
ALTER TABLE transactions ADD FOREIGN KEY (transaction_status_id) REFERENCES transaction_statuses (id);
