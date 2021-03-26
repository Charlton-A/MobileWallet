CREATE DATABASE mobilewallet;

CREATE SCHEMA  mobile;


CREATE TABLE mobile.users(
        id SERIAL PRIMARY KEY ,
        first_name VARCHAR(40) NOT NULL,
        last_name VARCHAR(50) NOT NULL,
        email VARCHAR(50) NOT NULL,
        phone VARCHAR(20) NULL,
        created_at TIMESTAMP DEFAULT  current_timestamp,
        updated_at TIMESTAMP NULL,
        deleted_at TIMESTAMP NULL,
        UNIQUE(email),
        UNIQUE(phone)

);


CREATE TABLE mobile.wallets(
        id SERIAL PRIMARY KEY ,
        user_id  INTEGER,
        balance  FLOAT DEFAULT 0.00 CHECK(balance >= 0.00),
        currency VARCHAR(10) DEFAULT 'SGD',
        created_at TIMESTAMP DEFAULT  current_timestamp,
        updated_at TIMESTAMP NULL,
        deleted_at TIMESTAMP NULL,
        CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES mobile.users(id) ON DELETE CASCADE
);


CREATE TABLE mobile.transaction_status(
        id SERIAL PRIMARY KEY,
        name VARCHAR(20),
        description  VARCHAR(100),
        created_at TIMESTAMP DEFAULT  current_timestamp,
        updated_at TIMESTAMP NULL,
        deleted_at TIMESTAMP NULL
);

CREATE TABLE mobile.transaction_type(
        id SERIAL PRIMARY KEY,
        name VARCHAR(20),
        description  VARCHAR(80),
        created_at TIMESTAMP DEFAULT  current_timestamp,
        updated_at TIMESTAMP NULL,
        deleted_at TIMESTAMP NULL
);


CREATE TABLE mobile.transactions(
        id SERIAL PRIMARY KEY ,
        trans_key VARCHAR(50),
        user_id  INTEGER,
        wallet_id  INTEGER,
        amount FLOAT NOT NULL,
        t_status INTEGER,
        t_type INTEGER,
        currency VARCHAR(10)  DEFAULT 'SGD',
        created_at TIMESTAMP DEFAULT  current_timestamp,
        updated_at TIMESTAMP NULL,
        deleted_at TIMESTAMP NULL,
        CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES mobile.users(id) ON DELETE CASCADE,
        CONSTRAINT fk_wallet FOREIGN KEY(wallet_id) REFERENCES mobile.wallets(id) ON DELETE CASCADE,
        CONSTRAINT fk_status FOREIGN KEY(t_status) REFERENCES mobile.transaction_status(id) ON DELETE CASCADE,
        CONSTRAINT fk_type FOREIGN KEY(t_type) REFERENCES mobile.transaction_type(id) ON DELETE CASCADE
);






INSERT INTO mobile.transaction_type(name,description) VALUES ('DEPOSIT', 'Deposit made to wallet'),
                                                        ('TRANSFER-DEBIT' ,'Funds transfered from  wallet') ,
                                                        ('TRANSFER-CREDIT' ,'Funds transfered to wallet'),
                                                        ('WITHDRAWAL', 'Withdrawal made from wallet');


INSERT INTO mobile.transaction_status(name,description) VALUES ('SUCCESS', 'Transaction was succesful'),
                                                        ('FAILED' ,'Transaction failed');









