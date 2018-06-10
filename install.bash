#!/bin/bash

# Automatically build and install serverManager2 for your system.


if ! type git >/dev/bull 2>&1; then
    echo "[ERR] Git needs to be installed to build this tool!"
    exit
fi

if ! type go >/dev/bull 2>&1; then
    echo "[ERR] Go needs to be installed to build this tool!"
    exit
fi

git clone https://github.com/zekroTJA/serverManager2.git src/github.com/zekroTJA/serverManager2
export GOAPTH=$PWD
echo "Cloning dependencies..."
go get github.com/logrusorgru/aurora
echo "Building and installing..."
go build -o /usr/bin/servermanager src/github.com/zekroTJA/serverManager2/main.go
echo "Cleaning up..."
rm -r pkg/ src/
echo "Installation finished: /usr/bin/servermanager"