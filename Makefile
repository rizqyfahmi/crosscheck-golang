.PHONY: app-build app-start app-start-daemon app-stop app-drop compose-up compose-down compose-clean

app-build:
	@-echo "Building image..."
	@docker build -t app-image . --progress=plain
	@-echo "Image's built successfully!"

app-start:
	@-echo "Running container..."
	@docker run -p 8081:8081 --name app-container app-image
	@-echo "Container's running successfully!"

app-start-daemon:
	@-echo "Running container in backrground..."
	@docker run -dp 8081:8081 --name app-container app-image
	@-echo "Container's running in backrground successfully!"

app-stop:
	@-echo "Stopping container..."
	@docker stop app-container
	@-echo "Dropping container..."
	@docker container rm app-container
	@-echo "Container's dropped successfully!"

app-drop: 
	@-echo "Dropping image..."
	@docker rmi app-image -f
	@-echo "Image's dropped successfully!"

# make compose-up
# make compose-up mode="daemon"
compose-up:
	@chmod -R 777 script
	@./script/compose-up.sh $(mode)

# make compose-down
# make compose-down mode="clean"
compose-down:
	@chmod -R 777 script
	@./script/compose-down.sh $(mode)

# make migrate-up
# make migrate-up version=1
migrate-up:
	@-echo "Migrating up..."
	@migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/crosscheck?sslmode=disable" -verbose up $(version)
	@-echo "Migrating up is success!"

# make migrate-down
# make migrate-down version=1
migrate-down:
	@-echo "Migrating down..."
	@migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/crosscheck?sslmode=disable" -verbose down $(version)
	@-echo "Migrating down is success!"

# make migrate-create name="create_table_profile"
migrate-create:
	@-echo "Creating migration file..."
	@migrate create -ext sql -dir migrations $(name)
	@-echo "Migration file successfully created..."

# make test-bootstrap path="tests"
test-bootstrap:
	@-echo "Creating test suite file..."
	@cd $(path) && ginkgo bootstrap
	@-echo "Test suite file created successfully..."

# make test-generate path="tests" name="utils"
test-generate:
	@-echo "Creating test suite file..."
	@cd $(path) && ginkgo generate $(name)
	@-echo "Test suite file created successfully..."
