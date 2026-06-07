//go:build country_all || country_asia || country_np || country_southern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNepal.RegisterName(xlanguage.English, "Nepal")
	dataNepal.RegisterOfficialName(xlanguage.English, "Federal Democratic Republic of Nepal")
	dataNepal.RegisterCapital(xlanguage.English, "Kathmandu")
}
