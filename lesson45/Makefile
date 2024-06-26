CURRENT_DIR=$(shell pwd)
DBURL := postgres://postgres:03212164@localhost:5432/library?sslmode=disable


proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}


mig-up:
	migrate -path database/migrations -database '${DBURL}' -verbose up

mig-down:
	migrate -path database/migrations -database '${DBURL}' -verbose down

mig-force:
	migrate -path database/migrations -database '${DBURL}' -verbose force 1



mig-create-books:
	migrate create -ext sql -dir database/migrations -seq create_books_table

mig-create-users:
	migrate create -ext sql -dir database/migrations -seq create_users_table

mig-create-rentals:
	migrate create -ext sql -dir database/migrations -seq create_rentals_table


# mig-delete:
#   rm -r db/migrations