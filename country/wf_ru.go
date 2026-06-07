//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWallisAndFutuna.RegisterName(xlanguage.Russian, "Уоллис и Футуна")
	dataWallisAndFutuna.RegisterOfficialName(xlanguage.Russian, "Территория островов Уоллис и Футуна")
	dataWallisAndFutuna.RegisterCapital(xlanguage.Russian, "Мата-Уту")
}
