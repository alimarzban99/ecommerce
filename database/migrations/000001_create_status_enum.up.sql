DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status') THEN
CREATE TYPE status AS ENUM ('active', 'inactive', 'blocked');
END IF;
END $$;