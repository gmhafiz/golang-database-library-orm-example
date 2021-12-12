#  Introduction

Examples of using various popular database libraries and ORM in Go.

 - [sqlx](https://jmoiron.github.io/sqlx/)
 - [sqlc](https://docs.sqlc.dev)
 - [Gorm](https://github.com/go-gorm/gorm)
 - [sqlboiler](https://github.com/volatiletech/sqlboiler)
 - [ent](https://entgo.io/docs/getting-started)

The aim is to demonstrate and compare usage for several operations

 1. Simple CRUD operation
 2. 1-1 queries
 3. 1-to-Many queries
 4. Many-to-many queries
 5. Dynamic list filter from query parameter 
 6. Transaction

The schema contains optional fields, for example middle name, and a field that must not be returned, for example, a password.


# Usage

## Setup

Setup postgres database by either running from docker-compose or manually.

    docker-compose up 

This creates both `postgres` database (which this repo uses) and `ent` database which is used by ent ORM.

Default database credentials are defined in `config/config.go`. These can be overwritten by setting environment variables. For example:

    export DB_NAME=test_db

## Run

Run with

    go run main.go


Run examples from `examples` folder.