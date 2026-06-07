//go:build (lang_fr || lang_all) && (country_all || country_nu || country_oceania || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiue.RegisterName(xlanguage.French, "Niue")
	dataNiue.RegisterOfficialName(xlanguage.French, "Niue")
	dataNiue.RegisterCapital(xlanguage.French, "Alofi")
}
