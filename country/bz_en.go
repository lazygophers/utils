//go:build country_all || country_americas || country_bz || country_central_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelize.RegisterName(xlanguage.English, "Belize")
	dataBelize.RegisterOfficialName(xlanguage.English, "Belize")
	dataBelize.RegisterCapital(xlanguage.English, "Belmopan")
}
