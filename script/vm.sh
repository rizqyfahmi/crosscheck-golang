main() {
    if [[ $(command docker images -q postgres) == '' ]]; then 
        echo "Pulling image..."
        command docker pull postgres:11-alpine
    fi

    if [[ $(command docker ps -aq -f name=app-database) == "" ]]; then 
        echo "Running postgres container..."
        command docker run --name app-database -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -v pwd:/var/lib/postgresql/data -d postgres:11-alpine
    fi


    command docker exec -it app-database psql -U postgres -c "SELECT 1 FROM pg_database WHERE datname = 'crosscheck'" | grep -q 1 || docker exec -it app-database psql -U postgres -c "CREATE DATABASE crosscheck"
    
}

main