//go:build (lang_ru || lang_all) && (country_all || country_asia || country_bn || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrunei.RegisterName(xlanguage.Russian, "Бруней")
	dataBrunei.RegisterOfficialName(xlanguage.Russian, "Государство Бруней-Даруссалам")
	dataBrunei.RegisterCapital(xlanguage.Russian, "Бандар-Сери-Бегаван")
}
