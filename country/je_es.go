//go:build (lang_es || lang_all) && (country_all || country_europe || country_je || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJersey.RegisterName(xlanguage.Spanish, "Jersey")
	dataJersey.RegisterOfficialName(xlanguage.Spanish, "Bailía de Jersey")
	dataJersey.RegisterCapital(xlanguage.Spanish, "Saint Helier")
}
