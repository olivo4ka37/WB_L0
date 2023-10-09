.PHONY:
.DEFAULT_GOAL := run
build:
		go mod download
		go build -o  ./.bin/app cmd/main.go

run:	build compose migrate
		./.bin/app

migrate:
		migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5040/l0db?sslmode=disable' up

dropTables:
		migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5040/l0db?sslmode=disable' down		

compose:
		docker-compose up -d

publish:
		go build -o ./.bin/publisher cmd/publisher/pub.go
		./.bin/publisher
