#!/bin/bash

echo > tmp/time
for i in {1..10}; do
	# care gtime (MacOs) = time (Linux)
    (gtime --f '%e' python py/genData_clone.py) 2>> tmp/time
done