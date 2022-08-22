.PHONY: app-build app-start app-start-daemon app-stop app-drop

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

.PHONY: compose-up compose-down
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

.PHONY: migrate-up migrate-down migrate-create
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

.PHONY: test-bootstrap test-generate test-generate-mock test-generate-mock-all test-run
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

# make test-generate-mock source="path/file.go" destination="path/file_mock.go" package="some_package"
test-generate-mock:
	@-echo "Generating mock file..."
	@mockgen -source=$(source) -destination=$(destination) -package=$(package)
	@-echo "Mock file successfully generated..."
# make test-generate-mock-all
test-generate-mock-all:
	@-echo "Generating all mock files..."
	@make test-generate-mock source="app/features/authentication/data/source/persistent/auth_persistent.go" destination="test/mocks/auth_persistent_mock.go" package="mock"
	@make test-generate-mock source="app/features/authentication/domain/repository/auth_repository.go" destination="test/mocks/auth_repository_mock.go" package="mock"
	@make test-generate-mock source="app/features/authentication/domain/usecase/registration/registration_usecase.go" destination="test/mocks/registration_usecase_mock.go" package="mock"
	@make test-generate-mock source="app/features/authentication/domain/usecase/login/login_usecase.go" destination="test/mocks/login_usecase_mock.go" package="mock"
	@make test-generate-mock source="app/utils/bcrypt/bcrypt.go" destination="test/mocks/bcrypt_mock.go" package="mock"
	@make test-generate-mock source="app/utils/clock/clock.go" destination="test/mocks/clock_mock.go" package="mock"
	@make test-generate-mock source="app/utils/hash/hash.go" destination="test/mocks/hash_mock.go" package="mock"
	@make test-generate-mock source="app/utils/jwt/jwt.go" destination="test/mocks/jwt_mock.go" package="mock"
	@-echo "All mock files successfully generated..."

# make test-run
test-run:
	@-echo "Running all test suites..."
	@ginkgo test ./...
	@-echo "All test suites successfully runned..."

.PHONY: http-registration
# make http-registration
http-registration:
	@-echo "Trying registration..."
	@curl -X POST http://localhost:8081/auth/registration -H "Content-Type: application/json" -d 'name=Rizqy Fahmi&email=rizqyfahmi@email.com&password=HelloPassword&confirmPassword=HelloPassword'
	@-echo "Registration successfully tried..."


.PHONY: doc-generate-spec doc-run
# make doc-generate-spec
doc-generate-spec: 
	@-echo "Generating swagger spec..."
	@swagger generate spec -o ./docs/swagger.yml --scan-models
	@-echo "Swagger spec successfully generated..."

# make doc-run
doc-run:
	@-echo "Running swagger spec..."
	@swagger serve -F=swagger -p 8082 ./docs/swagger.yml
	@-echo "Swagger spec successfully run..."
.PHONY: check
# make check
check:
	@-echo "Checking environment..."
	@sh ./script/doctor.sh
	@-echo "Environment checked successfully..."
