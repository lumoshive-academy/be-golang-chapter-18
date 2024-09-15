-- Membuat tabel users
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100),
    birth_date DATE,
    registration_date TIMESTAMP
);
