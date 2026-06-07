//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_eastern_africa || country_tz)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTanzania.RegisterName(xlanguage.MustParse("zh-Hant"), "坦尚尼亞")
	dataTanzania.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "坦尚尼亞聯合共和國")
	dataTanzania.RegisterCapital(xlanguage.MustParse("zh-Hant"), "杜篤瑪")
}
