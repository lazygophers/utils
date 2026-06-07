//go:build (lang_es || lang_all) && (country_all || country_americas || country_bz || country_central_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelize.RegisterName(xlanguage.Spanish, "Belice")
	dataBelize.RegisterOfficialName(xlanguage.Spanish, "Belice")
	dataBelize.RegisterCapital(xlanguage.Spanish, "Belmopán")
}
