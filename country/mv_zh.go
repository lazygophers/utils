//go:build country_all || country_asia || country_mv || country_southern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMaldives.RegisterName(xlanguage.Chinese, "马尔代夫")
	dataMaldives.RegisterOfficialName(xlanguage.Chinese, "马尔代夫共和国")
	dataMaldives.RegisterCapital(xlanguage.Chinese, "马累")
}
