-- Tabel products untuk menyimpan informasi produk
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    stock INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);
 
-- Tabel orders untuk menyimpan informasi pesanan
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
 
-- Tabel order_items untuk menyimpan detail item dalam pesanan
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
 
-- Mengisi data awal untuk tabel products
INSERT INTO products (name, stock, price) VALUES 
('Product A', 10, 50.00),
('Product B', 20, 30.00),
('Product C', 15, 20.00);

