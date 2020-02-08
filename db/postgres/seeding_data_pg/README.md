## Creating an application database

Log into postgres, then:

```sql
create database myapp;
create user myappuser with password 'asdf1234';
grant all privileges on database myapp to myappuser;
```

## "Migrations"

### What are migrations?

Changing a database is called "migrating" it.

Notice the `.sql` files in this directory, prefixed "000001", "000002", etc. 

The ascending numbers represent an ordered series of changes, or "migrations", made to the database.

There's always a pair of such changes -- "up" and "down".  

1. The "up" changes moves forward in time. 
1. The "down" change reverses the "up" change, thus moving backward in time. 

### Running the migrations

To use the other postgres/go snippets, you will need to run the "up" migrations in order. (See "run multiple sql commands from a file", below).

## Accessing the database from the command-line

### Run a single SQL command

```sh
# list the tables in the db
PGPASSWORD=asdf1234 psql -h localhost -U myappuser -d myapp -c "\dt"
```

### Run multiple SQL commands from a file

```sh
# run the 1st migration
PGPASSWORD=asdf1234 psql -h localhost -U myappuser -d myapp -f 000001_create_account_table.up.sql
```
