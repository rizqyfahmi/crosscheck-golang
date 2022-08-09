## Prolog

This aim of this project was creating a REST API that is implementing Clean Architecture, and TDD based on Golang. However, it has not covered possibilities that it also can be used for more reasons

## Prerequisites

1. Go 1.17
2. PostgreSQL 11
3. Git
4. Make
5. Docker

## Get Started

```sh
# Build and run the application
make compose-up

# Stop and drop/remove the application
make compose-down
```

## More Commands

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