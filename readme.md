# Setup Backend

## Running Database Migrations
To run migrations, I'm using Goose which is a library that manages data migration.

See installation for more detail:
https://github.com/pressly/goose?tab=readme-ov-file#install

Then after goose is installed, run this:
`goose -dir ./backend/data-migration sqlite3 vouchers.db up`

Then you should've seen a vouchers.db created in your project directory

## Running App
After running the migration, you can just run:
`cd backend` then after that `go run main.go`

# Setup Frontend

# Notes

## .env
I know it is not a good practice to push .env as it might contain password, where
usually a good practice is to make .env.local and .env to separate local setup and prod.

However for the sake of setup simplicity, the .env will be pushed regardless. 

## Validating Data
I'm using go validator in order to validate the request body that was sent from the frontend:
https://pkg.go.dev/github.com/go-playground/validator/v10

One of the use case is to validate the date must be in valid format (YYYY-MM-DD)
