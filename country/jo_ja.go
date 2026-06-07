//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJordan.RegisterName(xlanguage.Japanese, "ヨルダン")
	dataJordan.RegisterOfficialName(xlanguage.Japanese, "ヨルダン・ハシミテ王国")
	dataJordan.RegisterCapital(xlanguage.Japanese, "アンマン")
}
