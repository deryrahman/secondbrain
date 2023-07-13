# SecondBrain

Seamless note taking for ideas, quotes, moments, and anything with minimal features which only essential for storing, recalling, and organizing.

This project is still under development, anything related to the development process are written in `/docs` RFC.

## Dev Setup

**Setup for development**
```sh
DB_USER=<your_db_user>
DB_PASSWORD=<your_db_password>

make dev-tools
psql -U $DB_USER -c 'CREATE DATABASE secondbrain_dev;'
make db-migrate-up
```

**Run the server**
```sh
make build
./build/secondhand serve --dsn "postgresql://${DB_USER}:${DB_PASSWORD}@localhost:5432/secondbrain_dev?sslmode=disable"
```
