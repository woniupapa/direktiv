#!/usr/bin/env bash

RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color
NL=$'\n'

DIREKTION=$HOME/go/src/github.com/vorteil/direktion/build/direktion
DIREKCLI=$HOME/go/src/github.com/vorteil/direktiv/build/direkcli
DIREKTIV_API="http://localhost:80"
NAMESPACE=test

suite_dir=${1%/}

echo "Searching for tests in: $suite_dir"

tests=`find $suite_dir -maxdepth 1 -type d | tail --lines=+2 | sort`
count=`echo "$tests" | wc -l`
failed=0

compile_script () {
	script=$1

	wf=`$DIREKTION compile $script 2>/dev/null`
	status=$?
	if [ $status -ne 0 ]; then 
		echo -e "    ${RED}error${NC} encounted compiling script; run the following command to reproduce:";
		echo "      $DIREKTION compile $script";
		return $status
	fi

	# TODO: fix create/update cli side?
	tmp="/tmp/test-workflow.yaml"
	echo "$wf" > $tmp
	$DIREKCLI --url=$DIREKTIV_API workflows create $NAMESPACE $tmp
	status=$?
	rm $tmp
	if [ $status -ne 0 ]; then return $status; fi

}

compile_scripts () {
	test=$1
	scripts=`find $test -maxdepth 1 -type f -name '*.direktion'`
	sc=`echo "$scripts" | wc -l`
	
	idx=0
	for script in $scripts; do 
		let idx++
		echo "  compiling script ($idx/$sc): $script..."
		compile_script $script 
		status=$?
		if [ $status -ne 0 ]; then return $status; fi
	done

}

wipe_namespace () {
	# TODO: fix cli to output machine-readable stuff for scripting
	return 0
}

perform_test () {
	test=$1

	if [ ! -d "$test" ]; then echo -e "  ${RED}error${NC} no directory for test: $test"; return 1; fi

	wipe_namespace

	compile_scripts $test
	status=$?
	if [ $status -ne 0 ]; then return $status; fi

	# TODO: fix direkcli to support the wait query param
	$DIREKCLI --url=$DIREKTIV_API workflows execute $NAMESPACE "test"
	status=$?
	if [ $status -ne 0 ]; then return $status; fi

}

report_results () {
	status=$1
	if [ $status -eq 0 ]; then 
		echo -e "  ${GREEN}SUCCESS${NC}"
	else 
		echo -e "  ${RED}FAIL${NC}"
		let failed++
	fi
}

perform_tests () {
	i=0
	for test in $tests; do 
		let i++
		echo -e "${BLUE}Running test${NC} ($i/$count): $test..."
		perform_test $test 
		status=$?
		report_results $status
	done
}

individuals=0

perform_individual_tests () {
	shift
	for test in "$@"
	do
		let individuals++
		echo -e "${BLUE}Running test${NC} $individuals: $test..."
		perform_test "$suite_dir/$test"
		status=$?
		report_results $status
	done
}

init_namespace () {

	# TODO: fix existing namespace error on cli-side?
	# $DIREKCLI --url=$DIREKTIV_API namespaces create $NAMESPACE
	# status=$?
	# if [ $status -ne 0 ]; then return $status; fi

	return 0

}

init_namespace
status=$?
if [ $status -ne 0 ]; then exit $status; fi

perform_individual_tests

if [ $individuals -eq 0 ]; then 
	perform_tests
fi

if [ $failed -eq 0 ]; then 
	echo "All tests completed successfully."
	exit 0
else
	echo "$failed tests failed."
	exit 1
fi


