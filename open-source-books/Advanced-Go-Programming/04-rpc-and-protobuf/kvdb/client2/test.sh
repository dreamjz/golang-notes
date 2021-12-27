#!/usr/bin/env sh
for ((i=6;i<12;i++))
do
	go run ./main.go set "k-${i}" v1$i
done
