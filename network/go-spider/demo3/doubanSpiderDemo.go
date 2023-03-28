package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	//输入页数
	var start, end int

	fmt.Println("请输入起始页:")
	fmt.Scan(&start)
	fmt.Println("请输入尾页:")
	fmt.Scan(&end)

	//处理任务
	doWork(start, end)
}

func doWork(start int, end int) {
	channel := make(chan int, end-start+1)
	//根据页面遍历远程爬取页面
	for i := start; i <= end; i++ {
		go crawlDB(i, channel)

	}
	//监听任务处理结束状态
	for i := start; i <= end; i++ {
		fmt.Printf("第%d页已经完成\n", <-channel)
	}
}

func crawlDB(page int, channel chan int) {

	//远程请求豆瓣网页，并获取返回值
	url := "https://movie.douban.com/top250?start=" + strconv.Itoa((page-1)*25) + "&filter="
	time.Sleep(1 * time.Second)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err:", err)
		return
	}

	//生成文件
	fileName := "第" + strconv.Itoa(page) + "页电影榜.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	//解析返回结果
	fmt.Println("body:", result)
	//电影名
	titleRegex := regexp.MustCompile("<img width=\"100\" alt=\"(.*?)\" src=\"(.*?)\" class=\"\">")
	titles := titleRegex.FindAllStringSubmatch(result, -1)

	scoreRegex := regexp.MustCompile("<span class=\"rating_num\" property=\"v:average\">(.*?)</span>")
	scores := scoreRegex.FindAllStringSubmatch(result, -1)

	pplCountRegex := regexp.MustCompile("<span>(.*?)人评价</span>")
	pplCount := pplCountRegex.FindAllStringSubmatch(result, -1)

	file.WriteString("Tile|score|pplCount")
	fmt.Println("title count:", len(titles))
	for i := 0; i < len(titles); i++ {
		file.WriteString(titles[i][1] + "|" + scores[i][1] + "|" + pplCount[i][1] + "\n")
	}
	file.Close()
	//输出任务结束状态
	channel <- page

}

func HttpGet(url string) (result string, err error) {
	fmt.Println("url:", url)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	//避免豆瓣反爬虫，加入user-agent
	req.Header.Add("User-Agent", "myClient")
	resp, httpErr := client.Do(req)
	if httpErr != nil {
		fmt.Println("http.Get err:", err)
		err = httpErr
		return
	}
	defer resp.Body.Close()
	//读取响应体
	buf := make([]byte, 4098)
	for {

		n, readErr := resp.Body.Read(buf)
		fmt.Println("n:", n)
		if n == 0 {
			break
		}
		if readErr != nil && readErr != io.EOF {
			fmt.Println("请求网页异常:", readErr)
			err = readErr
			break
		}
		result += string(buf[:n])
		fmt.Println("result:", result)
	}
	return result, err
}
