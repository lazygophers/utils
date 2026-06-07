//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPuertoRico.RegisterName(xlanguage.Russian, "Пуэрто-Рико")
	dataPuertoRico.RegisterOfficialName(xlanguage.Russian, "Содружество Пуэрто-Рико")
	dataPuertoRico.RegisterCapital(xlanguage.Russian, "Сан-Хуан")
}
