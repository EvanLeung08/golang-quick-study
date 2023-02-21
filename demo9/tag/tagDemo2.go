package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name  string   `json:"name"`
	Price float32  `json:"rmb"`
	Type  []string `json:"type"`
}

func main() {
	movie := Movie{"星球大战", 16.1, []string{"恐怖", "爱情"}}

	//把对象编码成json格式
	marshal, err := json.Marshal(&movie)
	if err != nil {
		fmt.Println("Json解释异常")
		return
	}
	fmt.Printf("marshal=%s", marshal)

	//把Json字符串反编码

	error := json.Unmarshal(marshal, &movie)
	if error != nil {
		fmt.Println("反编码失败")
		return
	}

	fmt.Printf("打印对象 %v", movie)

}
