#!/bin/bash

if [ ! -e main ]; then
    echo 'Compiling...'
    go build main.go Tree.go
    if [ $? -eq 0 ]; then
        echo 'Compilation successful. Running program'
    fi
fi

./main

if [ ! "$1" = "c" ]; then
    rm main
fi