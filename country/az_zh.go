//go:build country_all || country_asia || country_az || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAzerbaijan.RegisterName(xlanguage.Chinese, "阿塞拜疆")
	dataAzerbaijan.RegisterOfficialName(xlanguage.Chinese, "阿塞拜疆共和国")
	dataAzerbaijan.RegisterCapital(xlanguage.Chinese, "巴库")
}
