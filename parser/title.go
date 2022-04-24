package parser

import (
	"chinadaily_com_cn/pkg/utils"
)

// Title 用于解析页面中的标题
func Title(body []byte) string {
	return utils.GetBetweenStr(string(body), "title: '", "',")
}
