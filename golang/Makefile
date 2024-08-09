MAKEFILE := $(lastword $(MAKEFILE_LIST))
SRC_DIRS := ./repository ./cmd ./internal/services ./internal/dtos \
            ./domain ./schema ./config ./utils ./api/v1/middlewares ./api/v1/routes \
			./internal/services/auth ./internal/services/blog ./internal/services/comment ./internal/services/reply \
			./internal/services/user ./internal/interfaces/repository_interface ./internal/interfaces/service_interfaces \
			 ./internal/dtos/auth  ./internal/dtos/blog  ./internal/dtos/user


.DEFAULT_GOAL := help


# ==================================================================================== #
# Help
# ==================================================================================== #

.PHONY: help
help:
	@echo 'Usage:'
	@echo
	@awk '/^##/{gsub("##", ""); print $$0}' $(MAKEFILE) | column -t -s ':' |  sed -e 's/^/ /'
	@echo


# ==================================================================================== #
# Formatting
# ==================================================================================== #


## format: uses to format files
.PHONY : format
format:
	@for dir in $(SRC_DIRS); do \
		go fmt $$dir/*.go; \
	done



# ==================================================================================== #
# Run
# ==================================================================================== #

## run : uses to run the project
.PHONY: run
run:
	go run cmd/main.go


## build-docker : uses to buld the docker image of the project
.PHONY: build-docker
build-docker:
	docker build -t golang .

## run-docker : uses to run the docker image of the project
.PHONY: run-docker
run-docker:
	docker run -it \
		--network app-network \
		-e PORT=8080 \
		-e DB_USER="root" \
		-e DB_PASSWORD="12345678" \
		-e DB_HOST="mysql" \
		-e DB_PORT="3306" \
		-e DB_NAME="golang" \
		-e JWT_SECRET_KEY="secretkey" \
		-e REDIS_DB=0 \
		-e REDIS_ADDR="redis:6379" \
		-e REDIS_PASS="12345678" \
		-p 8080:8080 \
		--entrypoint /bin/sh \
    golang 
	