
.PHONY: migrate_create

migrate_create:
	migrate create -ext sql -dir mysql/migrations -seq create_users_table
migrate:
	migrate -database 'mysql://root:password@tcp(localhost:3306)/careconnect' -path mysql/migrations up
