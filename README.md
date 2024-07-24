# Wallet-Wise
Simple self-hosted budgeting app

## Vue 3 + TypeScript + Vite

This template should help get you started developing with Vue 3 and TypeScript in Vite. The template uses Vue 3 `<script setup>` SFCs, check out the [script setup docs](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup) to learn more.

Learn more about the recommended Project Setup and IDE Support in the [Vue Docs TypeScript Guide](https://vuejs.org/guide/typescript/overview.html#project-setup).


## Requirements

Go, gcc

## Run

```shell
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3
```

Add to end of .bashrc
```shell
export PATH=$PATH:$HOME/go/bin
```

```shell
npm install
```

```shell
npm run build 
```

```shell
go run .
```

## CGO ERROR

Error:
```shell
Failed to open database: failed to ping database: Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
exit status 1

```

Fix:
```shell
go env -w CGO_ENABLED=1
```
