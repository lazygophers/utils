//go:build lang_ru || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrunei.RegisterName(xlanguage.Russian, "Бруней")
	dataBrunei.RegisterOfficialName(xlanguage.Russian, "Государство Бруней-Даруссалам")
	dataBrunei.RegisterCapital(xlanguage.Russian, "Бандар-Сери-Бегаван")
}
