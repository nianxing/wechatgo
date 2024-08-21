package main

import (
	"fmt"
  
	wxsg "github.com/nianxing/wechatgocrawler"
)

func main() {
	results, err := wxsg.SearchArticle("星阑文旅", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, result := range results {
		fmt.Printf("[%d] %s (%s) %s\n", i+1, result.Title, result.AccName, result.PubTime.String())
		fmt.Println(result.Preview)
		fmt.Println(result.Url)
	}
}
