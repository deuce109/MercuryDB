#!/bin/sh

# Directory to store test files
output_dir="../testFiles"

# Create the directory if it doesn't exist
mkdir -p "$output_dir"

# Generate 11 test files with 100 bytes of random data
for i in $(seq 1 11); do
    output_file="$output_dir/test_file_$i.txt"
    head -c 100 /dev/urandom > "$output_file"
    echo "Generated $output_file"
done

echo "All test files generated successfully."