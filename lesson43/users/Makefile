DB_URL=postgres://postgres:03212164@localhost:5432/users_service

run:
	go run cmd/server.go

migrate-file:
	migrate create -ext sql -dir database/migrations -seq users

migrate-up:
	migrate -path database/migrations -database ${DB_URL} -verbose up 

migrate-down:
	migrate -path database/migrations -database ${DB_URL} -verbose down 

migrate-force:
	migrate -path database/migrations -database ${DB_URL} -verbose force 1