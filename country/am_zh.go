//go:build country_all || country_am || country_asia || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataArmenia.RegisterName(xlanguage.Chinese, "亚美尼亚")
	dataArmenia.RegisterOfficialName(xlanguage.Chinese, "亚美尼亚共和国")
	dataArmenia.RegisterCapital(xlanguage.Chinese, "埃里温")
}
