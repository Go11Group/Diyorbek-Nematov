-- Users jadvali 

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL
);

-- Product jadvali

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INT NOT NULL
);

CREATE TABLE user_products (
    id SERIAL PRIMARY KEY,
    user_id int REFERENCES users(id),
    product_id int REFERENCES products(id)
)

-- Users jadvaliga 5 ta ma'lumot qo'shish
INSERT INTO users (username, email, password) 
VALUES 
    ('Alice', 'alice@example.com', 'password123'),
    ('Bob', 'bob@example.com', 'password456'),
    ('Charlie', 'charlie@example.com', 'password789'),
    ('David', 'david@example.com', 'password012'),
    ('Eve', 'eve@example.com', 'password345');

-- Products jadvaliga 5 ta ma'lumot qo'shish
INSERT INTO products (name, description, price, stock_quantity) 
VALUES 
    ('Laptop', 'High-performance laptop for professional use', 1299.99, 20),
    ('Headphones', 'Noise-canceling headphones for immersive sound experience', 199.99, 30),
    ('Smartwatch', 'Fitness tracker with smartwatch features', 149.99, 40),
    ('Camera', 'Digital camera with advanced photography capabilities', 899.99, 15),
    ('Tablet', 'Tablet device for entertainment and productivity', 499.99, 25);


-- user_products jadvaliga 5 ta ma'lumot qo'shish
INSERT INTO user_products (user_id, product_id) 
VALUES 
    (1, 1), -- Alice dan Laptop
    (2, 2), -- Bob dan Headphones
    (3, 3), -- Charlie dan Smartwatch
    (4, 4), -- David dan Camera
    (5, 5); -- Eve dan Tablet
