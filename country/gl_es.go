//go:build (lang_es || lang_all) && (country_all || country_americas || country_gl || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGreenland.RegisterName(xlanguage.Spanish, "Groenlandia")
	dataGreenland.RegisterOfficialName(xlanguage.Spanish, "Groenlandia")
	dataGreenland.RegisterCapital(xlanguage.Spanish, "Nuuk")
}
