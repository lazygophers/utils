//go:build country_all || country_micronesia || country_nr || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNauru.RegisterName(xlanguage.English, "Nauru")
	dataNauru.RegisterOfficialName(xlanguage.English, "Republic of Nauru")
	dataNauru.RegisterCapital(xlanguage.English, "Yaren")
}
