//go:build (lang_ko || lang_all) && (country_all || country_oceania || country_polynesia || country_tk)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTokelau.RegisterName(xlanguage.Korean, "토켈라우")
	dataTokelau.RegisterOfficialName(xlanguage.Korean, "토켈라우")
}
