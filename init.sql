CREATE DATABASE IF NOT EXISTS rest_api;

USE rest_api;

CREATE TABLE IF NOT EXISTS companies (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(50) NOT NULL,
    country VARCHAR(100) NOT NULL,
    website TEXT NOT NULL,
    phone VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY ( id )
);

INSERT INTO rest_api.companies (name, code, country, website, phone, created_at, updated_at) VALUES ('Apple', '1f22d', 'USA', 'https://apple.com', '+761209121233', DEFAULT, DEFAULT)