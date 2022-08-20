main() {
    if [[ $(command docker images -q postgres) == '' ]]; then 
        echo "Pulling image..."
        command docker pull postgres:11-alpine
    fi

    if [[ $(docker ps -aq -f name=app-database) != "" ]]; then 
        echo "Running postgres container..."
        command docker run --name app-database -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -v /data:/var/lib/postgresql/data -d postgres
    fi

}

main