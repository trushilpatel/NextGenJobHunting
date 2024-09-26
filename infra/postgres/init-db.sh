#!/bin/sh
set -e

echo "PostgreSQL is ready. Initializing database..."

# Create the PostgreSQL user and database from environment variables
psql -v ON_ERROR_STOP=1 --username ${POSTGRES_USER} --dbname "postgres" <<-EOSQL
    CREATE USER ${DB_USERNAME} WITH PASSWORD '${DB_PASSWORD}';
    CREATE DATABASE ${DB_NAME};
    GRANT ALL PRIVILEGES ON DATABASE ${DB_NAME} TO ${DB_USERNAME};

    \c ${DB_NAME} ${POSTGRES_USER}
    GRANT ALL ON SCHEMA public to ${DB_USERNAME};
EOSQL

echo "Database and user created successfully."
