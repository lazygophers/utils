//go:build (lang_ru || lang_all) && (country_all || country_americas || country_caribbean || country_pr)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPuertoRico.RegisterName(xlanguage.Russian, "Пуэрто-Рико")
	dataPuertoRico.RegisterOfficialName(xlanguage.Russian, "Содружество Пуэрто-Рико")
	dataPuertoRico.RegisterCapital(xlanguage.Russian, "Сан-Хуан")
}
