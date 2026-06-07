//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintHelena.RegisterName(xlanguage.Russian, "Острова Святой Елены, Вознесения и Тристан-да-Кунья")
	dataSaintHelena.RegisterOfficialName(xlanguage.Russian, "Острова Святой Елены, Вознесения и Тристан-да-Кунья")
	dataSaintHelena.RegisterCapital(xlanguage.Russian, "Джеймстаун")
}
