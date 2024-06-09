.PHONY: start
start:
		go run ./cmd/app/main.go

.PHONY: swag init
swag init:
			swag init -g internal/controller/http/v1/router.go

.PHONY: migrate-up
migrate-up:
			migrate -path migrations -database 'postgresql://Marat:1234@localhost/Test?sslmode=disable' up

.PHONY: migrate-down
migrate-down:
			migrate -path migrations -database 'postgresql://Marat:1234@localhost/Test?sslmode=disable' down
.PHONY: migrate-create
migrate-create:
				migrate create -ext sql -dir migrations 'migrate_name'




