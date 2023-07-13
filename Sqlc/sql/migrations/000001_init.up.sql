CREATE TABLE products (
    id VARCHAR(64) PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    description text
);

CREATE TABLE colors (
    id VARCHAR(64) PRIMARY KEY NOT NULL,
    product_id VARCHAR(64) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description text,
    price DECIMAL(10,2) NOT NULL,
    priceFinal DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id)
);
