DROP TABLE IF EXISTS users;
CREATE TABLE users ( 
    username TEXT NOT NULL PRIMARY KEY,
    password_hash TEXT NOT NULL, 
    created_date TIMESTAMP DEFAULT NOW(),
);
-- Ideally would salt the password as well
-- Pre create some users, 
-- TO DO: Need to have end point to create users
INSERT INTO users (username, password_hash) VALUES ('user1', 'e38ad214943daad1d64c102faec29de4afe9da3d'); 
INSERT INTO users (username, password_hash) VALUES ('user2', '2aa60a8ff7fcd473d321e0146afd9e26df395147'); 

DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts (
    id SERIAL PRIMARY KEY, 
    name TEXT NOT NULL,
    description TEXT, 
    created_date TIMESTAMP DEFAULT NOW(),
    owned_by TEXT NOT NULL
);

-- Pre create some accounts, 
-- TO DO: Need to have end point to create accounts
INSERT INTO accounts (name, description, owned_by) VALUES ('one','one','user1'); 
INSERT INTO accounts (name, description, owned_by) VALUES ('two','two','user2'); 


DROP TABLE IF EXISTS linked_customers;
CREATE TABLE linked_customers (
    customer_id TEXT NOT NULL PRIMARY KEY,
    account_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    total_revenue integer DEFAULT 0 NOT NULL
);


DROP TABLE IF EXISTS event_log;
CREATE TABLE event_log (
    event_id TEXT NOT NULL,
    event_type TEXT NOT NULL,
    payload TEXT NOT NULL,
    created_date TIMESTAMP DEFAULT NOW()
);

-- Have to set archival rules for event_log


-- Ideally should be in seperate databases, for now in the same one

DROP TABLE IF EXISTS customer;
CREATE TABLE customer (
    id SERIAL PRIMARY KEY, 
    account_id TEXT NOT NULL,
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




