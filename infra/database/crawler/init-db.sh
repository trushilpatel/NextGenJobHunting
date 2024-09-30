#!/bin/sh
set -e

echo "PostgreSQL is ready. Initializing database..."

# Create the PostgreSQL user and database from environment variables
psql -v ON_ERROR_STOP=1 --username ${POSTGRES_USER} --dbname "postgres" <<-EOSQL
    CREATE USER ${CRAWLER_DB_USERNAME} WITH PASSWORD '${CRAWLER_DB_PASSWORD}';
    CREATE DATABASE ${CRAWLER_DB_NAME};
    GRANT ALL PRIVILEGES ON DATABASE ${CRAWLER_DB_NAME} TO ${CRAWLER_DB_USERNAME};

    \c ${CRAWLER_DB_NAME} ${POSTGRES_USER}
    GRANT ALL ON SCHEMA public to ${CRAWLER_DB_USERNAME};
EOSQL

echo "##### Crawler Database and user created successfully."
