#!/bin/bash
make
./GinGen $(pwd)/functional_tests/file_test.go $(pwd)/functional_tests/output_test.json
diff $(pwd)/functional_tests/output_test.json $(pwd)/functional_tests/default_output.json
