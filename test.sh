#!/bin/sh
find . -name '*_test.go' | while read file; do
n=$(dirname -- "$file")
echo "$n"
done | sort -u | while read d; do
c=$(pwd)
cd "$d"
go test .
cd "$c"
done