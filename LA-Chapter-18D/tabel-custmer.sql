-- sql create tabel customers
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    email VARCHAR(100)
);

-- Insert multiple sample data
INSERT INTO customers (username, password, email) VALUES 
('admin', 'adminpassword', 'admin@example.com'),
('customer1', 'password1', 'customer1@example.com'),
('customer2', 'password2', 'customer2@example.com'),
('customer3', 'password3', 'customer3@example.com');
