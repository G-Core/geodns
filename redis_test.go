package main

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestRedis(t *testing.T) {

	pool := NewPool(":6379")

	conn := pool.Get()
	defer conn.Close()

	validConn := IsValidConn(conn)
	_ = validConn

	if true && validConn {
		res := RedisHGET(conn, "client", "155")

		fmt.Println(string(res))
		_ = res
	}

	if true && validConn {
		// BSON vs JSON vs ReJSON vs MSGPACK ...:
		// - JSON is most comfortable to deal with
		// - JSON speed is not so slow, maximum 3.5 than gob,
		//   http://danmux.com/posts/across_the_wire_serialisation/
		dat, err := json.Marshal(&ROPList{
			List: []string{
				"sv4-5",
				"se",
			},
			IsOut: true,
		})
		CheckError(err)

		reply, err := conn.Do("HSET", "client", "155", dat)
		CheckError(err)
		_, ok := reply.(int64)
		Assert(ok)

		var rl ROPList

		exists := GetROPList(conn, &rl, ROPClient{
			Key: "client",
			Value: "155",
		})
		Assert(exists)

		Assert(rl.List[0] == "sv4-5")
	}

	if true && validConn {
		dat := RedisHGET(conn, "client", "non-existent")
		Assert(dat == nil)
	}

	if true {
		dat1 := []byte{
			133,
		}

		dat2 := dat1[0:0]
		Assert(dat2 != nil)
		//fmt.Println(dat2)

		dat2 = nil
		Assert(dat2 == nil)
		//fmt.Println(string(dat2))

		var dat3 []string
		Assert(dat3 == nil)
	}
}
