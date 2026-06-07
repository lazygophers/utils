//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMoldova.RegisterName(xlanguage.MustParse("zh-Hant"), "摩爾多瓦")
	dataMoldova.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "摩爾多瓦共和國")
	dataMoldova.RegisterCapital(xlanguage.MustParse("zh-Hant"), "奇西瑙")
}
