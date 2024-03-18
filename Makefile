run:
	APP_ENV="production" PORT=5000 DB_USER=postgres DB_PASSWORD=postgres DB_NAME=hacktiv_deploy DB_HOST=localhost DB_PORT=5432 go run main.go 