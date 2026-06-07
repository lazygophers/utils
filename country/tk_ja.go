//go:build (lang_ja || lang_all) && (country_all || country_oceania || country_polynesia || country_tk)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTokelau.RegisterName(xlanguage.Japanese, "トケラウ")
	dataTokelau.RegisterOfficialName(xlanguage.Japanese, "トケラウ")
}
