.PHONY: run

run:
		go run cmd/card-project-api-server/main.go

build:
		docker-compose up --build -d
		# docker logs -f card-project | jq

swagger:
		swagger generate server -f ./swagger.yaml --exclude-main

test:
		go test -v -count=1 ./...
