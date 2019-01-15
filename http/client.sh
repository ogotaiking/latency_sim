#!/bin/bash
while true; do
	thread_number=$(ps -ef | grep client_linux | wc -l)
	while [ "$thread_number" -gt 3 ]; do
		sleep 1
		thread_number=$(ps -ef | grep client_linux | wc -l)
	done
	./client_linux -l 1 -n 50 -s 500 -t 1000 -u "101.0.1.1:8080" &
done
