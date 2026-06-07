//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelgium.RegisterName(xlanguage.Russian, "Бельгия")
	dataBelgium.RegisterOfficialName(xlanguage.Russian, "Королевство Бельгия")
	dataBelgium.RegisterCapital(xlanguage.Russian, "Брюссель")
}
