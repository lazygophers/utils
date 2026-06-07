//go:build country_all || country_europe || country_northern_europe || country_sj

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSvalbardAndJanMayen.RegisterName(xlanguage.Chinese, "斯瓦尔巴和扬马延")
	dataSvalbardAndJanMayen.RegisterOfficialName(xlanguage.Chinese, "斯瓦尔巴和扬马延")
	dataSvalbardAndJanMayen.RegisterCapital(xlanguage.Chinese, "朗伊尔城")
}
