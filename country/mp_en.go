//go:build country_all || country_micronesia || country_mp || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthernMarianaIslands.RegisterName(xlanguage.English, "Northern Mariana Islands")
	dataNorthernMarianaIslands.RegisterOfficialName(xlanguage.English, "Commonwealth of the Northern Mariana Islands")
	dataNorthernMarianaIslands.RegisterCapital(xlanguage.English, "Saipan")
}
