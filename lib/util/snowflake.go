package util

import (
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node
var epoch int64

func init() {
	t, _ := time.Parse("2006-01-02 15:04:05", "2018-01-01 00:00:00")
	epoch = t.UnixNano() / 1000000
	snowflake.Epoch = epoch
	nodeID := 1
	node, _ = snowflake.NewNode(int64(nodeID))
}

// SnowflakeID 创建一个唯一且顺序 int64 id
func NewSnowflakeID() int64 {
	return int64(node.Generate())
}

// SnowflakeID 创建一个唯一且顺序 string id
func NewSnowflakeIDBase58() string {
	return node.Generate().Base58()
}

// Int64ToBase58 转换一个id为短字符串形式
func Int64ToBase58(id int64) string {
	return snowflake.ID(id).Base58()
}

// Base58ToInt64 解析错误时返回 -1
func Base58ToInt64(s string) int64 {
	id, _ := snowflake.ParseBase58([]byte(s))
	return id.Int64()
}

// Int64ToBase36 转换一个id为短字符串形式
func Int64ToBase36(id int64) string {
	return strconv.FormatInt(id, 36)
}

// Base36ToInt64 解析错误时返回 -1
func Base36ToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 36, 64)
	return i
}
