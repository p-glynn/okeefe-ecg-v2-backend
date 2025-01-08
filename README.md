# This repository contains the backend server and database for the O'Keefe ECG v2 project.

## Database Setup

To use this database setup:

First, make sure you have PostgreSQL installed and running locally

1. Create a new database named ecg_db:

```bash
createdb ecg_db
```

1. Run the migration:

```bash
psql -d ecg_db -f db/migrations/001_initial_schema.sql
```

## Database Migrations

### Creating a New Migration

To create a new database migration:

```bash
migrate create -ext sql -dir db/migrations -seq <migration_name>
```

### Running Migrations

To apply all pending migrations:

```bash
migrate -path db/migrations -database "postgresql://username:password@localhost:5432/dbname?sslmode=disable" up
```

### Rolling Back Migrations

To roll back the last migration:

```bash
migrate -path db/migrations -database "postgresql://username:password@localhost:5432/dbname?sslmode=disable" down 1
```

Note: Replace `username`, `password`, `localhost`, `5432`, and `dbname` with your actual database connection details.
