#!/bin/bash

# Made this due to go template path won't be found by goland if run front/main.go in IDE.
# Unless you click the top of this bash, it then run from folder rather then tmp

echo "1. In goland, click green arrow above to run"

cd front

open http://localhost:3000/

echo "2. switch chrome to trigger refresh"

echo "3. the rest will block io"
go run main.go

