//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.Russian, "Мозамбик")
	dataMozambique.RegisterOfficialName(xlanguage.Russian, "Республика Мозамбик")
	dataMozambique.RegisterCapital(xlanguage.Russian, "Мапуту")
}
