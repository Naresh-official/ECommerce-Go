include .env
export

MIGRATIONS_DIR=migrations
MIGRATIONS_SOURCE=file://$(CURDIR)/$(MIGRATIONS_DIR)

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

migrate-up:
	migrate -source $(MIGRATIONS_SOURCE) -database "$(DATABASE_URL)" up

migrate-down-one:
	migrate -source $(MIGRATIONS_SOURCE) -database "$(DATABASE_URL)" down 1

migrate-down-all:
	migrate -source $(MIGRATIONS_SOURCE) -database "$(DATABASE_URL)" down -all

migrate-force:
	migrate -source $(MIGRATIONS_SOURCE) -database "$(DATABASE_URL)" force $(VERSION)

migrate-version:
	migrate -source $(MIGRATIONS_SOURCE) -database "$(DATABASE_URL)" version

sqlc:
	sqlc generate