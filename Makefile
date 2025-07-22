DB_URL=postgres://root:@localhost:5432/presensi_sekolah?sslmode=disable

up:
	migrate -database $(DB_URL) -path ./database/migrations up

drop:
	migrate -database $(DB_URL) -path ./database/migrations drop

docs: 
	swag init
	swag fmt

di:
	wire gen api/internal/injectors

docker-compose-up:
	docker compose -p presensi-sekolah up -d

docker-compose-down:
	docker compose down

.PHONY: up drop docs di docker-compose-up docker-compose-down
