#!/bin/bash

rm -r ./models

sqlboiler sqlite3 -c sqlboiler.toml --add-global-variants
