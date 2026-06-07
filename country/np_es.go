//go:build (lang_es || lang_all) && (country_all || country_asia || country_np || country_southern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNepal.RegisterName(xlanguage.Spanish, "Nepal")
	dataNepal.RegisterOfficialName(xlanguage.Spanish, "República Democrática Federal de Nepal")
	dataNepal.RegisterCapital(xlanguage.Spanish, "Katmandú")
}
