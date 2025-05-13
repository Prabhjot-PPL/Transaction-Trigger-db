-- To make a trigger do : 
-- psql -h localhost -p 5433 -U postgres -d e_commerce -f trigger.sql

-- To disable a trigger do : 
-- ALTER TABLE products DISABLE TRIGGER trg_set_updated_at;

-- Create the trigger function
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Drop existing trigger if it exists
DROP TRIGGER IF EXISTS trg_set_updated_at ON products;

-- Create the trigger``
CREATE TRIGGER trg_set_updated_at
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();