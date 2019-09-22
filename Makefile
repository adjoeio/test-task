build:
	./docker/build.sh

tests:
	@docker run --rm -v $(shell go env GOPATH):/go -v ${CURDIR}:/go/src/adjoe.io adjoe-test/golang-dev go test ./...
up:
	@make build
	docker-compose -f docker-compose.yml up -d
logs:
	docker-compose logs -f
rm:
	docker-compose rm  -sfv
start:
	docker-compose start
stop:
	docker-compose stop
rest:
	@make rm
	@make up
	@make logs

bash:
	docker-compose exec test-task bash
help:
	@echo "make container=adjoe_test-task_1 bash \t\t\t: exec a bash shell in the specific container"
