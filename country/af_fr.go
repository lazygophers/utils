//go:build (lang_fr || lang_all) && (country_af || country_all || country_asia || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAfghanistan.RegisterName(xlanguage.French, "Afghanistan")
	dataAfghanistan.RegisterOfficialName(xlanguage.French, "Émirat islamique d'Afghanistan")
	dataAfghanistan.RegisterCapital(xlanguage.French, "Kaboul")
}
