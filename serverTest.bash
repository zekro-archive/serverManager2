#!/bin/bash

# This script is for creating test envoirement with 100 test
# servers including an "application" script (in this case a simple
# script printing something to console every second) and a
# run script.
# After, the actuall test will be run

mkdir backups
mkdir -p testservers/test{1..100}
for ind in {1..100}
do
    echo "while true; do echo test123; sleep 1; done" >> testservers/test$ind/server.sh
    echo "cd \$1; bash server.sh" >> testservers/test$ind/run.sh
done

echo "SET UP TEST ENVOREMENT"
echo "STARTING TEST..."

go test

echo "TESTING COMMAND LINE ARGS"

echo "CREATING TEST CONFIG..."
echo '{' \
     '"serverLocation": "./testservers",' \
     '"backupLocation": "./backups",' \
     '"enableLogging": 1' \
     '}' > ./testconf.json

echo "HELP MESSAGE:"
./bin/build --test --help
echo "VERSION MESSAGE:"
./bin/build --test -v
echo "STARTING SERVERS 1, 2 AND 3:"
./bin/build --test -s test1,test2,test3
echo "STOPPING SERVERS 1, 2 AND 3:"
./bin/build --test -t test1,test2,test3
echo "STARTING SERVERS 1, 2 AND 3 IN LOOP MODE:"
./bin/build --test -s test1,test2,test3 --loop
echo "STOPPING SERVERS 1, 2 AND 3:"
./bin/build --test -t test1,test2,test3