DB_URL=postgresql://postgres:1421@localhost:5432/blog_db?sslmode=disable

swag-init:
	swag init -g api/api.go -o api/docs

start:
	go run cmd/main.go

migrateup:
		migrate -path migrations -database "$(DB_URL)" -verbose up

migratedown:
		migrate -path migrations -database "$(DB_URL)" -verbose down

.PHONY:	migrateup migratedown