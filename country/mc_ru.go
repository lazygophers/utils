//go:build (lang_ru || lang_all) && (country_all || country_europe || country_mc || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMonaco.RegisterName(xlanguage.Russian, "Монако")
	dataMonaco.RegisterOfficialName(xlanguage.Russian, "Княжество Монако")
	dataMonaco.RegisterCapital(xlanguage.Russian, "Монако")
}
