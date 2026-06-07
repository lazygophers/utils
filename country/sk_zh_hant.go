//go:build (lang_zh_hant || lang_all) && (country_all || country_eastern_europe || country_europe || country_sk)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovakia.RegisterName(xlanguage.MustParse("zh-Hant"), "斯洛伐克")
	dataSlovakia.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "斯洛伐克共和國")
	dataSlovakia.RegisterCapital(xlanguage.MustParse("zh-Hant"), "布拉提斯拉瓦")
}
