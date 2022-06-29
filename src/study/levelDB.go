package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	//"github.com/syndtr/goleveldb/leveldb/filter"
	//"github.com/syndtr/goleveldb/leveldb/opt"
)

func main() {
	db, err := leveldb.OpenFile("/Users/byunjaejin/Go/level_DB", nil)
	if err != nil {
		panic(err)
	}

	err = db.Put([]byte("0"), []byte("value00"), nil)
	err = db.Put([]byte("1"), []byte("value11"), nil)
	err = db.Put([]byte("2"), []byte("value22"), nil)

	data, err := db.Get([]byte("1"), nil)
	fmt.Println(decode(data))
}

// 바이트를 문자열로
func decode(b []byte) string {
	return string(b[:len(b)])
}
