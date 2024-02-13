package db

import (
	_ "github.com/mattn/go-sqlite3"
)

//go-sqlite3 package isnt directly used,but go uses it under the hood as we interact with the built in sql package part of go's std library. We append _ to it which tells go we need that import, although we dont use it directly, but it exposes functionality that is used under the hood by the built in sql package
