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