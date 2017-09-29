#!/bin/bash

echo > tmp/time

echo \#py1000 >> tmp/time
for i in {1..10}; do
	# care gtime (MacOs) = time (Linux)
    (gtime --f '%e' python py/client_clone.py 1000) 2>> tmp/time
done

echo \#go1000-1 >> tmp/time
go build go/client_clone.go
for i in {1..10}; do
	# care gtime (MacOs) = time (Linux)
    (gtime --f '%e' ./client_clone 1000 1) 2>> tmp/time
done
rm client_clone

echo \#go1000-1 >> tmp/time
for i in {1..10}; do
	# care gtime (MacOs) = time (Linux)
    (gtime --f '%e' go run go/client_clone.go 1000 1) 2>> tmp/time
done

echo \#go1000-2 >> tmp/time
for i in {1..10}; do
	# care gtime (MacOs) = time (Linux)
    (gtime --f '%e' go run go/client_clone.go 1000 2) 2>> tmp/time
done