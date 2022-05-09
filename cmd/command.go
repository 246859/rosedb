package main

import (
	"fmt"
	"github.com/flower-corp/rosedb"
	"github.com/tidwall/redcon"
)

const (
	resultOK = "OK"
)

func newWrongNumOfArgsError(cmd string) error {
	return fmt.Errorf("ERR wrong number of arguments for '%s' command", cmd)
}

func set(db *rosedb.RoseDB, args [][]byte) (interface{}, error) {
	if len(args) < 2 {
		return nil, newWrongNumOfArgsError("set")
	}
	key, value := args[0], args[1]
	err := db.Set(key, value)
	if err != nil {
		return nil, err
	}
	return redcon.SimpleString(resultOK), nil
}

func get(db *rosedb.RoseDB, args [][]byte) (interface{}, error) {
	if len(args) < 1 {
		return nil, newWrongNumOfArgsError("get")
	}
	value, err := db.Get(args[0])
	if err != nil {
		return nil, err
	}
	return value, nil
}

func keyType(db *rosedb.RoseDB, args [][]byte) (interface{}, error) {
	// todo
	return "string", nil
}

func info(db *rosedb.RoseDB, args [][]byte) (interface{}, error) {
	// todo
	return "info", nil
}