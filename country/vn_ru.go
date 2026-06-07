//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVietnam.RegisterName(xlanguage.Russian, "Вьетнам")
	dataVietnam.RegisterOfficialName(xlanguage.Russian, "Социалистическая Республика Вьетнам")
	dataVietnam.RegisterCapital(xlanguage.Russian, "Ханой")
}
