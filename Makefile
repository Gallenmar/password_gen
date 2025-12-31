help: ## Show help
	@echo "Usage: "
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

.PHONY: build run cleanup test clean

build: ## Build the container
	docker compose -f build/docker-compose.yml build

run: ## Run the password generator (usage: make run FLAGS="--length 30 --numbers")
	docker compose -f build/docker-compose.yml run --rm app go run ./cmd/password_gen/main.go $(FLAGS)

example: ## Run with example flags
	make run FLAGS="--length 10 --numbers --lower" --no-print-directory

cleanup: ## Delete history
	docker compose -f build/docker-compose.yml run --rm app go run ./cmd/password_gen/main.go --cleanup

test: ## Run tests
	docker compose -f build/docker-compose.yml run --rm app go test -v ./...

clean: ## Clean up resources
	make cleanup --no-print-directory
	docker compose -f build/docker-compose.yml down --volumes --remove-orphans