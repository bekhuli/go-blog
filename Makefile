MIGRATE=go run cmd/migrate/main.go

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

migrate-drop:
	$(MIGRATE) drop

migration:
	@if [ -z "$(word 2, $(MAKECMDGOALS))" ]; then \
		echo "Usage: make new-migration <name>"; \
	else \
		migrate create -ext sql -dir migrations -seq $(word 2, $(MAKECMDGOALS)); \
	fi

%:
	@:
