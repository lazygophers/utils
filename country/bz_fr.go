//go:build (lang_fr || lang_all) && (country_all || country_americas || country_bz || country_central_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelize.RegisterName(xlanguage.French, "Belize")
	dataBelize.RegisterOfficialName(xlanguage.French, "Belize")
	dataBelize.RegisterCapital(xlanguage.French, "Belmopan")
}
