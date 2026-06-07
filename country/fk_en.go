//go:build country_all || country_americas || country_fk || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFalklandIslands.RegisterName(xlanguage.English, "Falkland Islands")
	dataFalklandIslands.RegisterOfficialName(xlanguage.English, "Falkland Islands")
	dataFalklandIslands.RegisterCapital(xlanguage.English, "Stanley")
}
