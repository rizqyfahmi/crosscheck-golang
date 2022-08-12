#!/usr/bin/env bash
main() {
    if [[ $(command docker images -q app-image) = '' ]]; then 
        echo "Building image"
        command docker-compose -p crosscheck build
    fi

    if [[ $1 = "daemon" ]]; then
        echo "Execute in daemon mode"
        command docker-compose up -d
    else
        echo "Execute in normal mode"
        command docker-compose up
    fi
}

main $1