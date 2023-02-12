#!/bin/bash
./GinGen $(pwd)/functional_tests/testing_file $(pwd)/functional_tests/output_test.json
diff $(pwd)/functional_tests/output_test.json $(pwd)/functional_tests/default_output.json
