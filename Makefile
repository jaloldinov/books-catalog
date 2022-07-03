migrateup:
	migrate --path migrations/postgres/ --database "postgresql://postgres:Muhammad@localhost:5432/books_shop?sslmode=disable" -verbose up

migratedown:
	migrate --path migrations/postgres/ --database "postgresql://postgres:Muhammad@localhost:5432/books_shop?sslmode=disable" -verbose down

run: 
	go run api/main.go

swag-init:
	swag init -g api/main.go -o api/docs

go-run:
	go run api/main.go

