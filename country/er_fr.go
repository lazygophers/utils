//go:build (lang_fr || lang_all) && (country_africa || country_all || country_eastern_africa || country_er)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEritrea.RegisterName(xlanguage.French, "Érythrée")
	dataEritrea.RegisterOfficialName(xlanguage.French, "État d'Érythrée")
	dataEritrea.RegisterCapital(xlanguage.French, "Asmara")
}
