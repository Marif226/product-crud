CREATE TABLE IF NOT EXISTS buyers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    contact VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS purchases (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    description TEXT,
    quantity INT NOT NULL,
    price INT NOT NULl,
    buyer_id INT REFERENCES buyers (id) ON DELETE CASCADE NOT NULL
);