#!/bin/bash

if [[ -z "$1" ]]; then
  echo "Usage: $0 <aoc_year>"
  exit 1
fi

baseDir="$1"

for i in {01..25};
do
	inner_dir="$baseDir/day-$i"
	mkdir -p $inner_dir
    export aoc_year="${baseDir}"
    export aoc_day=$(expr $i + 0)
	envsubst < "scripts/solution_template.txt" > "${inner_dir}/solution.go"
	envsubst < "scripts/solution_test_template.txt" > "${inner_dir}/solution_test.go"

	mkdir "${inner_dir}/testdata"
done
tree "${baseDir}"
