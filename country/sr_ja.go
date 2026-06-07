//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSuriname.RegisterName(xlanguage.Japanese, "スリナム")
	dataSuriname.RegisterOfficialName(xlanguage.Japanese, "スリナム共和国")
	dataSuriname.RegisterCapital(xlanguage.Japanese, "パラマリボ")
}
