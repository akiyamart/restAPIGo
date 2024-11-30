migrate_base_up: 
	migrate -path platform/migrations -database "postgres://postgres:postgres@localhost:5434/postgres?sslmode=disable" up
migrate_base_down: 
	migrate -path platform/migrations -database "postgres://postgres:postgres@localhost:5434/postgres?sslmode=disable" down