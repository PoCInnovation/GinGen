#!/bin/bash

GREEN="\033[1;32m"
CYAN="\033[1;36m"
NO_COLOR="\033[0m"

# Check arguments when executing scritp
if [ "$#" -ne 1 ]; then
    echo "USAGE: $0 <generator>"
    echo "EXAMPLE: $0 rust"
    exit 0
fi

make
./GinGen -i /demo/testing_file -o output_demo.json
echo -e "${GREEN}➜  ${CYAN}[Successfully generate json spec]${NO_COLOR}"
echo -e "${GREEN}➜  ${CYAN}[Start to generate file]${NO_COLOR}"
openapi-generator-cli generate  -i demo/output.json -g $1 -o server
