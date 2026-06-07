//go:build (lang_ru || lang_all) && (country_all || country_americas || country_central_america || country_gt)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuatemala.RegisterName(xlanguage.Russian, "Гватемала")
	dataGuatemala.RegisterOfficialName(xlanguage.Russian, "Республика Гватемала")
	dataGuatemala.RegisterCapital(xlanguage.Russian, "Гватемала")
}
