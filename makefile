TEST_COVERAGE_OUTPUT := test_coverage

test: FORCE
	go test -p 1 ./... -coverprofile=${TEST_COVERAGE_OUTPUT}

FORCE: ;

show-test-cover:
	go tool cover -html=${TEST_COVERAGE_OUTPUT}
