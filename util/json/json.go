package json

import (
	jsoniter "github.com/json-iterator/go"
)

// 假設有天需要替換 為了方便 以及美觀
// 不想要外面都加上 var json = jsoniter.ConfigCompatibleWithStandardLibrary 特地建這個包
// 外面統一都用 json._____ 呼叫  有需要什麼func 就在這加
// 要替換包的話直接把 import gorder/util/json 路徑改為自己要的

var j = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(v interface{}) ([]byte, error) {
	return j.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return j.Unmarshal(data, v)
}
