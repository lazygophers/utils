//go:build (lang_ru || lang_all) && (country_all || country_americas || country_central_america || country_pa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPanama.RegisterName(xlanguage.Russian, "Панама")
	dataPanama.RegisterOfficialName(xlanguage.Russian, "Республика Панама")
	dataPanama.RegisterCapital(xlanguage.Russian, "Панама")
}
