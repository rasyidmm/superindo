setup:
	@echo " --- Setup and generate configuration --- "
	@cp internal/config/example/database.yml.example internal/config/db/database.yml
	@cp internal/config/example/server.yml.example internal/config/server/server.yml
	@cp internal/config/example/redis.yml.example internal/config/redis/redis.yml
	swag init  #init swagger
	@echo " --- Done Setup and generate configuration --- "

rest:
	@go run main.go

build: setup
	@echo "--- Building binary file ---"
	@go build -o ./main main.go