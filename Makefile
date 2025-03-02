# Makefile

# Variables
N = 10

clean:
	docker-compose -f docker-compose.yml down

deploy: clean build
	docker-compose -f docker-compose.yml up --build --remove-orphans
test:
	@for i in $(shell seq 1 $(N)); do \
  		curl -X POST -H "Content-Type: application/json" -d '{"input": {"role": "admin", "password":"somepass", "state" :"TX"}}' http://localhost:8181/v1/data/authz/allow; \
  		curl -X POST -H "Content-Type: application/json" -d '{"input": {"role": "user","password":"somepass" , "state" :"TX"}}' http://localhost:8181/v1/data/authz/allow; \
	done
build:
	@rm -fR ./package/bundle.tar.gz
	@opa build bundle
	@mv bundle.tar.gz ./package/
	@tar -tzf ./package/bundle.tar.gz


clean70:
	docker-compose -f docker-compose70.yml down

deploy70: clean70 build70
	docker-compose -f docker-compose70.yml up --build --remove-orphans

build70:
	@rm -fR ./package/bundle.tar.gz
	@opa70 build bundle70
	@mv bundle.tar.gz ./package/
	@tar -tzf ./package/bundle.tar.gz