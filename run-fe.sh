#!/bin/bash

echo "1. In goland, click green arrow above to run"

cd front

open http://localhost:3000/

echo "2. switch chrome to trigger refresh"

echo "3. the rest will block io"
go run main.go

