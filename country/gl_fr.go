//go:build (lang_fr || lang_all) && (country_all || country_americas || country_gl || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreenland.RegisterName(xlanguage.French, "Groenland")
	dataGreenland.RegisterOfficialName(xlanguage.French, "Groenland")
	dataGreenland.RegisterCapital(xlanguage.French, "Nuuk")
}
