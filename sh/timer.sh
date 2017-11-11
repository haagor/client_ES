#!/bin/bash

LOOP=100
echo > tmp/time

IT=0
echo \#py1000 >> tmp/time
while [  $IT -lt $LOOP ]; do
	# care gtime (MacOs) = time (Linux)
    (gtime --f '%e' python py/client_clone.py 1000) 2>> tmp/time
    let IT=IT+1
done

IT=0
echo \#pypy1000 >> tmp/time
while [  $IT -lt $LOOP ]; do
	# care gtime (MacOs) = time (Linux)
    (gtime --f '%e' ../pypy2-v5.9.0-osx64/bin/pypy py/client_clone.py 1000) 2>> tmp/time
    let IT=IT+1
done

for j in {1..10}; do

	IT=0
	echo \#goB1000-$j >> tmp/time
	go build go/client_clone.go
	while [  $IT -lt $LOOP ]; do
		# care gtime (MacOs) = time (Linux)
    	(gtime --f '%e' ./client_clone 1000 $j) 2>> tmp/time
    	let IT=IT+1
	done
	rm client_clone

	IT=0
	echo \#go1000-$j >> tmp/time
	while [  $IT -lt $LOOP ]; do
		# care gtime (MacOs) = time (Linux)
    	(gtime --f '%e' go run go/client_clone.go 1000 $j) 2>> tmp/time
    	let IT=IT+1
	done

done

python py/averageTime.py $LOOP
