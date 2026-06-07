//go:build country_africa || country_all || country_eastern_africa || country_mg

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMadagascar.RegisterName(xlanguage.French, "Madagascar")
	dataMadagascar.RegisterOfficialName(xlanguage.French, "République de Madagascar")
	dataMadagascar.RegisterCapital(xlanguage.French, "Antananarivo")
}
