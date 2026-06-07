//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBurundi.RegisterName(xlanguage.Russian, "Бурунди")
	dataBurundi.RegisterOfficialName(xlanguage.Russian, "Республика Бурунди")
	dataBurundi.RegisterCapital(xlanguage.Russian, "Гитега")
}
