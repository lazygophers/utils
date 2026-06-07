//go:build country_all || country_asia || country_kw || country_western_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKuwait.RegisterName(xlanguage.English, "Kuwait")
	dataKuwait.RegisterOfficialName(xlanguage.English, "State of Kuwait")
	dataKuwait.RegisterCapital(xlanguage.English, "Kuwait City")
}
