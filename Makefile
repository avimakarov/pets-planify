prod.up:
	COMPOSE_ENV_FILES=./env/.prod docker-compose -f docker-compose.yml up --build -d

prod.down:
	COMPOSE_ENV_FILES=./env/.prod docker-compose -f docker-compose.yml down

generate.server:
	oapi-codegen -generate="types,chi-server" -package=server_gen ./etc/openapi/server/server.yaml > ./internal/generated/openapi/server/server.go

test.integration:
	-@echo "Starting integration tests..."
	-$(MAKE) test.integration.up
	-$(MAKE) test.integration.run
	-$(MAKE) test.integration.down
	-$(MAKE) test.integration.drop_data
	-@echo "Integration tests completed (with possible failures)"

test.integration.up:
	COMPOSE_ENV_FILES=./env/.test docker-compose -f docker-compose.yml up postgres postgres_migrations --build -d

test.integration.run:
	set -a; source ./env/.test; set +a; go test --tags=integration -count=1 -v ./internal/...

test.integration.down:
	COMPOSE_ENV_FILES=./env/.test docker-compose -f docker-compose.yml down

test.integration.drop_data:
	rm -R ./data