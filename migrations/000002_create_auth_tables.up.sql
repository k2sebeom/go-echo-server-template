CREATE TABLE IF NOT EXISTS password_credentials (
  id SERIAL PRIMARY KEY,

  user_id INTEGER UNIQUE REFERENCES users(id) ON DELETE CASCADE,

  email VARCHAR(30) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL,

  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

DROP TYPE IF EXISTS social_provider;
CREATE TYPE social_provider
AS ENUM ('KAKAO', 'NAVER', 'GOOGLE');

CREATE TABLE IF NOT EXISTS social_credentials (
  id SERIAL PRIMARY KEY,

  user_id INTEGER REFERENCES users(id),

  identifier_hash VARCHAR(255) NOT NULL,
  provider social_provider NOT NULL,

  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);