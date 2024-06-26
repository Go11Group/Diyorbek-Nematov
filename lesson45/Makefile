CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

exp:
	export DBURL='postgres://postgres:03212164@localhost:5432/library?sslmode=disable'

mig-up:
	migrate -path migrations -database 'postgres://postgres:03212164@localhost:5432/library?sslmode=disable' -verbose up

mig-down:
	migrate -path migrations -database ${DBURL} -verbose down


mig-create-books:
	migrate create -ext sql -dir database/migrations -seq create_books_table

mig-create-users:
	migrate create -ext sql -dir database/migrations -seq create_users_table

mig-create-rentals:
	migrate create -ext sql -dir database/migrations -seq create_rentals_table


# mig-delete:
#   rm -r db/migrations