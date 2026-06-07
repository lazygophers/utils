//go:build country_all || country_asia || country_south_eastern_asia || country_vn

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVietnam.RegisterName(xlanguage.English, "Vietnam")
	dataVietnam.RegisterOfficialName(xlanguage.English, "Socialist Republic of Vietnam")
	dataVietnam.RegisterCapital(xlanguage.English, "Hanoi")
}
