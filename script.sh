#!/bin/bash

run_command() {
    output=$(go run .)
    first_line=$(echo "$output" | head -n 1)
    if [ "$first_line" != "Found a collision" ]; then
        return 1
    fi
    echo "$output"
}

while ! run_command; do
    echo "Retrying..."
done
