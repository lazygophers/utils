//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataChina.RegisterName(xlanguage.Japanese, "中華人民共和国")
	dataChina.RegisterOfficialName(xlanguage.Japanese, "中華人民共和国")
	dataChina.RegisterCapital(xlanguage.Japanese, "北京")
}
