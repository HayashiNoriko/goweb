// 使用 Lua 脚本实现原子操作
package main

import (
	"fmt"
)

func main() {
	rdb, ctx := initRedis()
	defer rdb.Close()

	script := `
	local key = KEYS[1]
	local value = ARGV[1]

	return redis.call('SET', key, value)
	`

	sha, err := rdb.ScriptLoad(ctx, script).Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 执行脚本
	_, err = rdb.EvalSha(ctx, sha, []string{"luakey"}, "luavalue...").Result()
	if err != nil {
		fmt.Println(err)
		return
	}
}
