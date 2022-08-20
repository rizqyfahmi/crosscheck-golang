main() {

    command docker network create internal-container-network || command docker network ls

    if [[ $(command docker images -q postgres) == '' ]]; then 
        echo "Pulling image..."
        command docker pull postgres:11-alpine
    fi

    command docker network create internal-container-network || command docker network ls
    command docker run --name app-database -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres --network internal-container-network -p 5432:5432 -v pwd:/var/lib/postgresql/data -d postgres:11-alpine || command docker ps -a
    command docker exec app-database psql -U postgres -c "CREATE DATABASE crosscheck" || echo "Database is created"
}

main