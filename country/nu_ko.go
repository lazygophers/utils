//go:build (lang_ko || lang_all) && (country_all || country_nu || country_oceania || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiue.RegisterName(xlanguage.Korean, "니우에")
	dataNiue.RegisterOfficialName(xlanguage.Korean, "니우에")
	dataNiue.RegisterCapital(xlanguage.Korean, "알로피")
}
