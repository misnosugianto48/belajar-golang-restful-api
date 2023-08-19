-- Active: 1689174877887@@127.0.0.1@3306@belajar_golang_restful_api

CREATE TABLE
    category (
        category_id VARCHAR(255) NOT NULL,
        name VARCHAR(30) NOT NULL,
        created_at TIMESTAMP,
        updated_at TIMESTAMP,
        PRIMARY KEY(category_id)
    );