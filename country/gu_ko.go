//go:build (lang_ko || lang_all) && (country_all || country_gu || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuam.RegisterName(xlanguage.Korean, "괌")
	dataGuam.RegisterOfficialName(xlanguage.Korean, "괌")
	dataGuam.RegisterCapital(xlanguage.Korean, "하갓냐")
}
