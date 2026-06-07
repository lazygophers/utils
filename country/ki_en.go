//go:build country_all || country_ki || country_micronesia || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKiribati.RegisterName(xlanguage.English, "Kiribati")
	dataKiribati.RegisterOfficialName(xlanguage.English, "Republic of Kiribati")
	dataKiribati.RegisterCapital(xlanguage.English, "Tarawa")
}
