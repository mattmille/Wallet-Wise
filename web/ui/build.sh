#!/bin/bash

rm -rf ./../../src/views

npm install
npm run build

mkdir -p ./../../src/views
cp -rp ./build/* ./../../src/views
