//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPapuaNewGuinea.RegisterName(xlanguage.Russian, "Папуа — Новая Гвинея")
	dataPapuaNewGuinea.RegisterOfficialName(xlanguage.Russian, "Независимое Государство Папуа — Новая Гвинея")
	dataPapuaNewGuinea.RegisterCapital(xlanguage.Russian, "Порт-Морсби")
}
