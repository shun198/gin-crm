prepare:
	docker-compose up -d --build

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

migrate:
	docker-compose exec app go run github.com/steebchen/prisma-client-go migrate dev

seed:
	docker-compose exec app go run prisma/seed.go

