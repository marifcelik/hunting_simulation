#!/bin/bash

# Function to run the command and check its output
run_command() {
    output=$(go run .)
    first_line=$(echo "$output" | head -n 1)
    if [ "$first_line" != "Found a collision" ]; then
        return 1
    fi
    echo "$output"
}

# Run the command until the desired condition is met
while ! run_command; do
    echo "Retrying..."
done
