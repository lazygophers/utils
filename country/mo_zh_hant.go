//go:build country_all || country_asia || country_eastern_asia || country_mo

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMacao.RegisterName(xlanguage.MustParse("zh-Hant"), "澳門")
	dataMacao.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "中華人民共和國澳門特別行政區")
	dataMacao.RegisterCapital(xlanguage.MustParse("zh-Hant"), "澳門")
}
