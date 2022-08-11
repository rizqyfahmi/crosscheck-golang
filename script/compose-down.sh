main() {
    if [[ $1 = "clean" ]]; then
        stopContainer
        removeImage
        resetVolumeDirectory
    else
        stopContainer
    fi
}

resetVolumeDirectory() {
    if [[ -d "volume" ]]; then
        echo "Removing volume directory..."
        command rm -rf volume
        echo "Recreating volume directory..."
        command mkdir volume
        echo "Change mode of volume directory..."
        command chmod -R 777 volume
        echo "Volume directory successfully reset..."
    else 
        echo "Reset volumen directory...(Skipped)"
    fi
}

stopContainer() {
    if [[ $(docker ps -aq -f name=app-container) != "" ]] || [[ $(docker ps -aq -f name=app-database) != "" ]]; then
        echo "Stopping container..."
        command docker-compose down
    else 
        echo "Stopping container...(Skipped)"
    fi
}

removeImage() {
    if [[ $(command docker images -q app-image) != '' ]]; then 
        echo "Removing image..."
        command make app-drop
    else 
        echo "Removing image...(Skipped)"
    fi
}

main $1