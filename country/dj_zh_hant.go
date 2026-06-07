//go:build (lang_zh_hant || lang_all) && (country_africa || country_all || country_dj || country_eastern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDjibouti.RegisterName(xlanguage.MustParse("zh-Hant"), "吉布地")
	dataDjibouti.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "吉布地共和國")
	dataDjibouti.RegisterCapital(xlanguage.MustParse("zh-Hant"), "吉布地市")
}
