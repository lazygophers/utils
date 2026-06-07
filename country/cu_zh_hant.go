//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_caribbean || country_cu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuba.RegisterName(xlanguage.MustParse("zh-Hant"), "古巴")
	dataCuba.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "古巴共和國")
	dataCuba.RegisterCapital(xlanguage.MustParse("zh-Hant"), "哈瓦那")
}
