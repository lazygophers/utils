//go:build country_africa || country_all || country_eastern_africa || country_ke

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKenya.RegisterName(xlanguage.English, "Kenya")
	dataKenya.RegisterOfficialName(xlanguage.English, "Republic of Kenya")
	dataKenya.RegisterCapital(xlanguage.English, "Nairobi")
}
