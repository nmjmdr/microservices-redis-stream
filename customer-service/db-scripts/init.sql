DROP TABLE IF EXISTS customer;
CREATE TABLE customer (
    id SERIAL PRIMARY KEY, 
    name TEXT NOT NULL, 
    created_date TIMESTAMP DEFAULT NOW()
);

DROP TABLE IF EXISTS invoice;
CREATE TABLE invoice (
    id SERIAL PRIMARY KEY, 
    customer_id INTEGER REFERENCES customer(id) ON DELETE CASCADE,
    description TEXT, 
    created_date TIMESTAMP DEFAULT NOW(), 
    purchase_date TIMESTAMP NOT NULL, 
    purchase_price DECIMAL NOT NULL
);

