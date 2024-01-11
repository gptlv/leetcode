#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Usage: $0 <folder_number>"
    exit 1
fi

folder_number=$1
folder_name="$folder_number"

# Create folder
mkdir "$folder_name"
cd "$folder_name" || exit 1

# Execute go mod init
go mod init "$folder_name"

# Create main.go with a basic main function
echo "package main

func main() {

}" > main.go

echo "Project setup completed in folder: $folder_name"
