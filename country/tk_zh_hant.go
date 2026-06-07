//go:build (lang_zh_hant || lang_all) && (country_all || country_oceania || country_polynesia || country_tk)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTokelau.RegisterName(xlanguage.MustParse("zh-Hant"), "托克勞")
	dataTokelau.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "托克勞")
}
