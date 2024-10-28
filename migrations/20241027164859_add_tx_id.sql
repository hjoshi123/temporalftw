ALTER TABLE transactions ADD COLUMN transaction_id UUID DEFAULT gen_random_uuid();
