//go:build country_africa || country_all || country_southern_africa || country_sz

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEswatini.RegisterName(xlanguage.Chinese, "斯威士兰")
	dataEswatini.RegisterOfficialName(xlanguage.Chinese, "斯威士兰王国")
	dataEswatini.RegisterCapital(xlanguage.Chinese, "姆巴巴内")
}
