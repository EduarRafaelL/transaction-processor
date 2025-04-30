-- Create base schema
CREATE TABLE clients (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);

CREATE TABLE  transaction_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    client_id VARCHAR(50) NOT NULL,
    date TIMESTAMP NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    transaction_type_id INT NOT NULL REFERENCES transaction_types(id),
    FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE TABLE templates (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    html_body TEXT NOT NULL
);


