CREATE TABLE categories (
    id INT PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE courses (
    id INT PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    category_id  int,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
