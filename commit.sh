#!/bin/bash

# Check if there are no arguments provided
if [ "$#" -eq 0 ]; then
    echo "Usage: $0 <problem_number> [<problem_number> ...]"
    exit 1
fi

# Iterate through each provided problem number
for problem in "$@"; do
    # Check if the folder for the problem number exists
    if [ -d "$problem" ]; then
        # Add the folder to the staging area
        git add "$problem"
        
        # Commit the changes with a meaningful message
        git commit -m "add: $problem solution"
        echo "Committed solution for problem $problem"
    else
        echo "Folder for problem $problem not found. Skipping."
    fi
done

# Push changes to the remote repository
git push
