//go:build country_all || country_nu || country_oceania || country_polynesia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiue.RegisterName(xlanguage.English, "Niue")
	dataNiue.RegisterOfficialName(xlanguage.English, "Niue")
	dataNiue.RegisterCapital(xlanguage.English, "Alofi")
}
