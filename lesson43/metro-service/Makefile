DB_URL=postgres://postgres:03212164@localhost:5432/metro_service

run:
	go run cmd/main.go

migrate-file:
	migrate create -ext sql -dir database/migrations -seq create_table

migrate-up:
	migrate -path database/migrations -database ${DB_URL} -verbose up 

migrate-down:
	migrate -path database/migrations -database ${DB_URL} -verbose down 

migrate-force:
	migrate -path database/migrations -database ${DB_URL} -verbose force 1