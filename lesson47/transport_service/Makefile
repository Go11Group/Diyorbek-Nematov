CURRENT_DIR=$(shell pwd)
DBURL := postgres://postgres:03212164@localhost:5432/transport?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}


mig-up:
	migrate -path database/migrations -database '${DBURL}' -verbose up

mig-down:
	migrate -path database/migrations -database '${DBURL}' -verbose down

mig-force:
	migrate -path database/migrations -database '${DBURL}' -verbose force 1


mig-create-transport:
	migrate create -ext sql -dir database/migrations -seq create_transport_info

mig-create-schedule:
	migrate create -ext sql -dir database/migrations -seq create_schedule