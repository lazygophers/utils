//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBosniaAndHerzegovina.RegisterName(xlanguage.Russian, "Босния и Герцеговина")
	dataBosniaAndHerzegovina.RegisterOfficialName(xlanguage.Russian, "Босния и Герцеговина")
	dataBosniaAndHerzegovina.RegisterCapital(xlanguage.Russian, "Сараево")
}
