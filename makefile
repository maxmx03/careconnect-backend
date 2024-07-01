
.PHONY: migrate_create, private_pem

migrate_install:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz
migrate_create:
	./migrate create -ext sql -dir mysql/migrations -seq create_user_table
	./migrate create -ext sql -dir mysql/migrations -seq create_doctor_table
	./migrate create -ext sql -dir mysql/migrations -seq create_patient_table
	./migrate create -ext sql -dir mysql/migrations -seq create_medicalprescription_table
	./migrate create -ext sql -dir mysql/migrations -seq create_message_table
	./migrate create -ext sql -dir mysql/migrations -seq create_medical_consultation_table
migrate_run:
	./migrate -database 'mysql://root:password@tcp(localhost:3307)/careconnect' -path mysql/migrations up
private_pem:
	openssl ecparam -genkey -name prime256v1 -out ec256-private.pem
