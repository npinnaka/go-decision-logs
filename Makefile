# Makefile

# Variables
N = 10

# Default target
all: clean deploy

clean:
	docker-compose down

deploy: clean build
	docker-compose up --build --remove-orphans
test:
	@for i in $(shell seq 1 $(N)); do \
  		curl -X POST -H "Content-Type: application/json" -d '{"input": {"role": "admin"}}' http://localhost:8181/v1/data/authz/allow; \
  		curl -X POST -H "Content-Type: application/json" -d '{"input": {"role": "user"}}' http://localhost:8181/v1/data/authz/allow; \
	done
build:
	@rm -fR ./package/bundle.tar.gz
	@opa build bundle
	@mv bundle.tar.gz ./package/
	@tar -tzf ./package/bundle.tar.gz
