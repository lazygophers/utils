//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahrain.RegisterName(xlanguage.Russian, "Бахрейн")
	dataBahrain.RegisterOfficialName(xlanguage.Russian, "Королевство Бахрейн")
	dataBahrain.RegisterCapital(xlanguage.Russian, "Манама")
}
