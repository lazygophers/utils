//go:build (lang_ru || lang_all) && (country_africa || country_all || country_sn || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSenegal.RegisterName(xlanguage.Russian, "Сенегал")
	dataSenegal.RegisterOfficialName(xlanguage.Russian, "Республика Сенегал")
	dataSenegal.RegisterCapital(xlanguage.Russian, "Дакар")
}
