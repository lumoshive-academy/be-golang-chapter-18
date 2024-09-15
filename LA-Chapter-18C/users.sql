-- Membuat tabel users (bahan praktik type data column)
CREATE TABLE users ( 
    id SERIAL PRIMARY KEY, 
    first_name VARCHAR(50), 
    last_name VARCHAR(50), 
    email VARCHAR(100), 
    birth_date DATE, 
    registration_date TIMESTAMP 
); 

--- membuat table users (bahan praktik null label)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    age INT,
    email VARCHAR(100),
    birth_date DATE,
    registration_date TIMESTAMP
);

--- menambahkan data pada table users (bahan praktik null label)
INSERT INTO users (username, age, email, birth_date, registration_date)
VALUES
    ('john_doe', 30, 'john.doe@example.com', '1990-01-15', NOW()),
    ('jane_smith', NULL, 'jane.smith@example.com', '1992-05-20', NOW()),
    (NULL, 25, 'james.johnson@example.com', '1997-09-10', NOW()),
    ('mary_brown', 28, NULL, '1993-11-25', NOW()),
    ('sam_roberts', 35, 'sam.roberts@example.com', NULL, NOW()),
    ('anna_green', 29, 'anna.green@example.com', '1995-04-30', NULL);
