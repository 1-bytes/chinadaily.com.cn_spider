package main

import (
	"chinadaily_com_cn/bootstrap"
	"chinadaily_com_cn/cmd"
	"chinadaily_com_cn/parser"
	"chinadaily_com_cn/pkg/config"
	"chinadaily_com_cn/pkg/fetcher"
	"chinadaily_com_cn/pkg/queued"
	"fmt"
	"github.com/gocolly/colly/v2"
	"regexp"
)

func main() {
	urls := []string{
		"https://www.chinadaily.com.cn/a/202204/22/WS6262aca8a310fd2b29e58c2c.html",
		"https://www.chinadaily.com.cn/a/202204/24/WS62650ee7a310fd2b29e58f44.html",
		"https://www.chinadaily.com.cn/a/202204/24/WS6265032aa310fd2b29e58f1f.html",
		"https://www.chinadaily.com.cn/a/202204/22/WS626214fda310fd2b29e58a04.html",
		"https://www.chinadaily.com.cn/a/202204/01/WS6246ad01a310fd2b29e54b28.html",
		"https://www.chinadaily.com.cn/a/202204/24/WS6264a764a310fd2b29e58e3b.html",
	}
	bootstrap.Setup()
	c := cmd.NewCollector(
		//colly.Debugger(&debug.LogDebugger{}),
		colly.Async(config.GetBool("spider.async", false)),
		colly.AllowedDomains("www.chinadaily.com.cn", "chinadaily.com.cn"),
		colly.DetectCharset(),
		colly.URLFilters(
			regexp.MustCompile(`www\.chinadaily\.com\.cn/a/\d{6}/\d{2}/[a-zA-Z0-9]+\.html?$`),
			//regexp.MustCompile(`http://www\.chinadaily\.com\.cn/index\.html$`),
		))
	cmd.SpiderCallbacks(c)

	for _, url := range urls {
		_ = queued.Queued.AddURL(url)
	}
	_ = queued.Queued.Run(c)
	//testCase()
}

// testCase 测试用例
func testCase() {
	bootstrap.Setup()
	url := "https://www.chinadaily.com.cn/a/202204/22/WS6262aca8a310fd2b29e58c2c.html"
	bytes, err := fetcher.Fetch(url)
	if err != nil {
		panic(err)
	}
	//id := parser.ID(url)
	title := parser.Title(bytes)
	author := parser.Author(bytes)
	category := parser.Category(url)
	releaseDate := parser.ReleaseDate(bytes)
	paragraphs, err := parser.Content(bytes)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	for _, paragraph := range paragraphs {
		//fmt.Printf("ID: %d\n", id)
		fmt.Printf("Title: %s\n", title)
		fmt.Printf("Author: %s\n", author)
		fmt.Printf("Category: %s\n", category)
		fmt.Printf("ReleaseDate: %s\n", releaseDate)
		fmt.Printf("EN: %s\n", paragraph["EN"])
		fmt.Printf("CN: %s\n", paragraph["CN"])
		fmt.Println()

		data := parser.JsonData{
			//ID:        strconv.Itoa(id),
			SourceURL: url,
			Paragraph: paragraph,
		}
		if err = cmd.SaveDataToElastic("chinadaily_com_cn", "", data); err != nil {
			fmt.Printf("SaveData error: %v\n", err)
		}
	}
}
