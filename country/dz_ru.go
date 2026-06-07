//go:build (lang_ru || lang_all) && (country_africa || country_all || country_dz || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlgeria.RegisterName(xlanguage.Russian, "Алжир")
	dataAlgeria.RegisterOfficialName(xlanguage.Russian, "Алжирская Народная Демократическая Республика")
	dataAlgeria.RegisterCapital(xlanguage.Russian, "Алжир")
}
