migrateup:
	migrate --path migrations/postgres/ --database "postgresql://postgres:Muhammad@localhost:5432/books_shop?sslmode=disable" -verbose up

migratedown:
	migrate --path migrations/postgres/ --database "postgresql://postgres:Muhammad@localhost:5432/books_shop?sslmode=disable" -verbose down

run: 
	go run api/main.go

swag-go:
	swag init -g api/api.go -o api/docs
	go run api/main.go

