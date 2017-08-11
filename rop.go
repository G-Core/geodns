package main

import (
	"github.com/garyburd/redigo/redis"
	"encoding/json"
)

type ROPList struct {
	List  []string `json:"list"`
	// is to exclude ROPs of the list
	IsOut bool `json:"is_out"`
}

func RedisHGET(conn redis.Conn, table string, key string) []byte {
	reply, err := conn.Do("HGET", table, key)
	CheckError(err)

	var res []byte
	if reply != nil {
		var ok bool
		res, ok = reply.([]byte)
		Assert(ok)
	}

	return res
}

type ROPClient struct {
	Key   string
	Value string
	FallbackLabel string
}

func GetROPList(conn redis.Conn, rl *ROPList, rc ROPClient) bool {
	dat := RedisHGET(conn, rc.Key, rc.Value)

	exists := dat != nil

	if exists {
		err := json.Unmarshal(dat, &rl)
		CheckError(err)
	}
	return exists
}

// no connection, for example
func IsValidConn(conn redis.Conn) bool {
	return conn.Err() == nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Assert(b bool)  {
	if !b {
		panic("Assertion error")
	}
}

