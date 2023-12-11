#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <directory_name>"
    exit 1
fi

dir_name="$1"

mkdir -p "$dir_name"

cd "$dir_name"

go mod init Advent-of-Code-23/$dir_name

touch main.go

touch input.txt

touch test.txt