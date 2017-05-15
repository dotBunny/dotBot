#!/bin/bash
cd "$(dirname "$0")"
go build -o dot ./src
./dot
