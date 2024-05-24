CREATE TYPE STATUS as ENUM('pending', 'completed', 'cancelled');

CREATE TABLE IF NOT EXISTS orders (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    total DECIMAL(10, 2) NOT NULL,
    status STATUS NOT NULL DEFAULT 'pending',
    address TEXT NOT NULL,
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);