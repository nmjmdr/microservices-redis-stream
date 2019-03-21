DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts (
    id SERIAL PRIMARY KEY, 
    name TEXT NOT NULL,
    description TEXT, 
    created_date TIMESTAMP DEFAULT NOW(),
    owned_by TEXT NOT NULL
);

DROP TABLE IF EXISTS linked_customers;
CREATE TABLE linked_customers (
    customer_id TEXT NOT NULL PRIMARY KEY,
    total_revenue NUMBER DEFAULT 0 NOT NULL
);


DROP TABLE IF EXISTS last_recorded_event_id;
CREATE TABLE last_recorded_event_id (
    last_event_id TEXT NOT NULL
);


