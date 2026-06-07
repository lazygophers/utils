//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSenegal.RegisterName(xlanguage.MustParse("zh-Hant"), "塞內加爾")
	dataSenegal.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "塞內加爾共和國")
	dataSenegal.RegisterCapital(xlanguage.MustParse("zh-Hant"), "達卡")
}
