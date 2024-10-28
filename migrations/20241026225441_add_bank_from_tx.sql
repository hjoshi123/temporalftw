ALTER TABLE transactions ADD COLUMN to_bank_id INT NOT NULL DEFAULT 1;
ALTER TABLE transactions ADD FOREIGN KEY (to_bank_id) REFERENCES banks (id);

ALTER TABLE transactions RENAME COLUMN bank_id TO from_bank_id;
