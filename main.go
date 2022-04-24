package main

import (
	"chinadaily_com_cn/bootstrap"
	"chinadaily_com_cn/cmd"
	"chinadaily_com_cn/parser"
	"chinadaily_com_cn/pkg/fetcher"
	"fmt"
	"github.com/gocolly/colly/v2"
	"regexp"
	"time"
)

func main() {
	bootstrap.Setup()
	c := cmd.NewCollector(
		//colly.Debugger(&debug.LogDebugger{}),
		colly.Async(true),
		colly.AllowedDomains("www.chinadaily.com.cn", "chinadaily.com.cn"),
		colly.DetectCharset(),
		colly.URLFilters(
			regexp.MustCompile(`www\.chinadaily\.com\.cn/\w*?/\d{6}/\d{2}/\w+\.html?$`),
			//regexp.MustCompile(`http://www\.chinadaily\.com\.cn/index\.html$`),
		))
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 20,
		RandomDelay: 200 * time.Millisecond,
	})
	cmd.SpiderCallbacks(c)
	if err := c.Visit("https://www.chinadaily.com.cn/a/202204/22/WS6262aca8a310fd2b29e58c2c.html"); err != nil {
		panic(err)
	}
	c.Wait()
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
