SHELL := /bin/bash
SCRIPT_DIRECTORY := ./scripts
MIGRATION_DIRECTORY := ./docs/migrations

.PHONY: migration-up
migration-up: migration-validate
	@$(SHELL) $(SCRIPT_DIRECTORY)/migration.sh "$(username)" "$(password)" "$(dbHost)" "$(databaseName)" "up" "$(qty)" "$(MIGRATION_DIRECTORY)"

.PHONY: migration-down
migration-down: migration-validate
	@$(SHELL) $(SCRIPT_DIRECTORY)/migration.sh "$(username)" "$(password)" "$(dbHost)" "$(databaseName)" "down" "$(qty)" "$(MIGRATION_DIRECTORY)"

.PHONY: migration-drop
migration-drop: migration-validate
	@$(SHELL) $(SCRIPT_DIRECTORY)/migration.sh "$(username)" "$(password)" "$(dbHost)" "$(databaseName)" "drop -f" "$(MIGRATION_DIRECTORY)"

.PHONY: migration-create
migration-create: migration-validate
	@$(SHELL) $(SCRIPT_DIRECTORY)/create_migration.sh "$(MIGRATION_DIRECTORY)" "$(migrationName)"

.PHONY: migration-validate
migration-validate:
	@$(SHELL) $(SCRIPT_DIRECTORY)/validate_migration.sh

.PHONY: swaggo
swaggo:
	@$(SHELL) $(SCRIPT_DIRECTORY)/validate_swaggo.sh
	@`go env GOPATH`/bin/swag fmt
	@`go env GOPATH`/bin/swag init

.PHONY: build
build: swaggo
	@go build -o build/main

.PHONY: run
run: build
	@./build/main