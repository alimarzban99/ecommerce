CREATE TABLE IF NOT EXISTS verification_codes (
    id SERIAL PRIMARY KEY,
    code VARCHAR(6) NOT NULL UNIQUE,
    mobile VARCHAR(15) NOT NULL,
    status status DEFAULT 'active',
    expire_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
 );

CREATE INDEX idx_verification_codes_mobile_code ON verification_codes (mobile, code);
