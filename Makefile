# Makefile

# Variables
N = 10

# Default target
all: clean deploy

clean:
	docker-compose down

deploy: clean
	docker-compose up --build --remove-orphans
test:
	@for i in $(shell seq 1 $(N)); do \
  		curl -XPOST -d '{ "input": { "user": "john", "password": "secret" } }' http://localhost:8181/v1/data/authz/allow; \
	done