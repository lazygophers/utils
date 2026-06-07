//go:build (lang_ja || lang_all) && (country_all || country_australia_and_new_zealand || country_nz || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewZealand.RegisterName(xlanguage.Japanese, "ニュージーランド")
	dataNewZealand.RegisterOfficialName(xlanguage.Japanese, "ニュージーランド")
	dataNewZealand.RegisterCapital(xlanguage.Japanese, "ウェリントン")
}
