//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_dm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominica.RegisterName(xlanguage.Russian, "Доминика")
	dataDominica.RegisterOfficialName(xlanguage.Russian, "Содружество Доминики")
	dataDominica.RegisterCapital(xlanguage.Russian, "Розо")
}
