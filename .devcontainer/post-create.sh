#!/bin/bash

# SQLBoiler stuff
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3
go build .

# Allows SQLBoiler binary to be found
export PATH=$PATH:$HOME/go/bin


# NPM stuff
npm install
npm run build


# Make sqlite3 db if it does not exist
if [ -f "./db.sqlite" ]; then
    echo "sqlite3 db already exists"
else
    sqlite3 db.sqlite < db.sql
fi
