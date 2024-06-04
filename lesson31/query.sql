CREATE TABLE Product (
    ID uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)

CREATE INDEX product_name ON product(name);
CREATE INDEX product_price_name ON product(price, name);
CREATE INDEX product_name_price ON product(name, price) ;