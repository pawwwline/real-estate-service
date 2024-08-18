CREATE TABLE house (
	id SERIAL PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    year INT NOT NULL CHECK (year >= 0),
    developer VARCHAR(255),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE flat (
    id SERIAL PRIMARY KEY,
    house_id INT NOT NULL,
    price INT NOT NULL CHECK (price >= 0),
    rooms INT NOT NULL CHECK (rooms >= 1),
    status VARCHAR(20) NOT NULL CHECK (status IN ('created', 'approved', 'declined', 'on moderation')),
    FOREIGN KEY (house_id) REFERENCES house(id) ON DELETE CASCADE
    );

CREATE INDEX idx_flat_house_id ON flat (house_id);

CREATE INDEX idx_status ON flat (status);