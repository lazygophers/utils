//go:build (lang_ru || lang_all) && (country_all || country_asia || country_ps || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalestine.RegisterName(xlanguage.Russian, "Палестина")
	dataPalestine.RegisterOfficialName(xlanguage.Russian, "Государство Палестина")
	dataPalestine.RegisterCapital(xlanguage.Russian, "Рамалла")
}
