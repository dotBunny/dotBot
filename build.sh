#!/bin/bash
cd "$(dirname "$0")"
go build -o bin/application ./src
