## Prolog

This aim of this project was creating a REST API that is implementing Clean Architecture, and TDD based on Golang. However, it has not covered possibilities that it also can be used for more reasons

## Prerequisites

1. Go 1.17
2. PostgreSQL 11
3. Git
4. Brew
5. Make
6. Docker
7. Golang-migrate
8. Ginkgo
9. Mockgen


## Get Started

```sh
# Check the prerequisites
make check

# Run the application and database container/image
make compose-up

# Build and run the application and database container/image in daemon mode
make compose-up mode="daemon"

# Stop the application and database container/image
make compose-down

# Stop and drop the application and database container/image
make compose-down mode="clean"
```

## Migration Commands

```sh
# Create a migration file
make migrate-create name="create_table"

# Apply all migrations
make migrate-up

# Apply some of migration files up to specific version
make migrate-up version=1

# Reverse all migrations
make migrate-down version=1

# Reverse some of migration files down to specific version
make migrate-down version=1
```

## Testing Commands

Make sure that you already have test suite on the same path you generate test file

```sh
# create a test suite into specific path
make test-bootstrap path="tests"

# create a test file into specific path
make test-generate path="tests" name="utils"

# create a mock file
make test-generate-mock source="path/file.go" destination="path/file_mock.go" package="some_package"

# running all test suites
make test-run
```

## More Commands

All commands below is used for managing application container independently (exclude database)
```sh
# Build the image of the application
make app-build

# Run the container of the application
make app-start

# Run the container of the application daemonly
make app-start-daemon

# Stop the container of the application
make app-stop

# Drop/remove the image of the application
make app-drop
```