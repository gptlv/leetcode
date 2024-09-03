#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Usage: $0 <folder_names>"
    exit 1
fi

folder_name=$1
folder="$folder_name"

# Create folder
mkdir "$folder"
cd "$folder" || exit 1

# Execute go mod init
go mod init "$folder"

# Create main.go with a basic main function
echo "package main

func main() {

}" > main.go

echo "Project setup completed in folder: $folder"
