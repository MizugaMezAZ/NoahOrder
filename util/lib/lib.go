package lib

import "fmt"

func RedisFormatKey(prefix string, key any) string {
	return fmt.Sprintf("%v-%v", prefix, key)
}
