include .env

gup:
	cd sql/schema && goose turso ${DBURL}?authToken=${DBTOKEN} up

gdown:
	cd sql/schema && goose turso ${DBURL}?authToken=${DBTOKEN} down

gstatus:
	cd sql/schema && goose turso ${DBURL}?authToken=${DBTOKEN} status

sqlc:
	sqlc generate

run:
	go run ./cmd/thisorthat/
