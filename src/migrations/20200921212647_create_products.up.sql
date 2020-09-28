CREATE TABLE products (
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255),
    article VARCHAR(255) NULL UNIQUE,
    price FLOAT NOT NULL,
    category_id BIGINT NULL,
    FOREIGN KEY (category_id) REFERENCES product_categories (id)
);