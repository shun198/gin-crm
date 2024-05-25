CONTAINER_NAME = app
RUN_APP = docker-compose exec $(CONTAINER_NAME)
RUN_TERRAFORM = docker-compose -f infra/docker-compose.yml run --rm terraform

prepare:
	docker-compose run --rm front npm install
	docker-compose up -d --build

up:
	docker-compose up -d

build:
	docker-compose build

down:
	docker-compose down

init:
	$(RUN_TERRAFORM) init

fmt:
	$(RUN_TERRAFORM) fmt

validate:
	$(RUN_TERRAFORM) validate

show:
	$(RUN_TERRAFORM) show

apply:
	$(RUN_TERRAFORM) apply -auto-approve

destroy:
	$(RUN_TERRAFORM) destroy