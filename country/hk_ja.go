//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHongKong.RegisterName(xlanguage.Japanese, "香港")
	dataHongKong.RegisterOfficialName(xlanguage.Japanese, "中華人民共和国香港特別行政区")
	dataHongKong.RegisterCapital(xlanguage.Japanese, "香港")
}
