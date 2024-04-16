
.PHONY: migrate_create

migrate_create:
	migrate create -ext sql -dir mysql/migrations -seq create_user_table
	migrate create -ext sql -dir mysql/migrations -seq create_doctor_table
	migrate create -ext sql -dir mysql/migrations -seq create_patient_table
	migrate create -ext sql -dir mysql/migrations -seq create_medicalprescription_table
	migrate create -ext sql -dir mysql/migrations -seq create_medication_table
	migrate create -ext sql -dir mysql/migrations -seq create_message_table
migrate:
	migrate -database 'mysql://root:password@tcp(localhost:3306)/careconnect' -path mysql/migrations up
