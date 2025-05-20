CREATE TABLE IF NOT EXISTS user_table (
    user_id UUID NOT NULL REFERENCES users(id),
    role_id UUID NOT NULL REFERENCES roles(id)
)