#!/bin/bash

# This script is for creating test envoirement with 100 test
# servers including an "application" script (in this case a simple
# script printing something to console every second) and a
# run script.
# After, the actuall test will be run

mkdir -p testservers/test{1..100}
for ind in {1..100}
do
    echo "while true; do echo test123; sleep 1; done" >> testservers/test$ind/server.sh
    echo "cd \$1; bash server.sh" >> testservers/test$ind/run.sh
done

echo "SET UP TEST ENVOREMENT"
echo "STARTING TEST..."

go test