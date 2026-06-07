//go:build (lang_zh_hant || lang_all) && (country_all || country_europe || country_northern_europe || country_sj)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSvalbardAndJanMayen.RegisterName(xlanguage.MustParse("zh-Hant"), "斯瓦巴及揚馬延")
	dataSvalbardAndJanMayen.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "斯瓦巴及揚馬延")
	dataSvalbardAndJanMayen.RegisterCapital(xlanguage.MustParse("zh-Hant"), "朗伊爾城")
}
