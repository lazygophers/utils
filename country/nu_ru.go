//go:build (lang_ru || lang_all) && (country_all || country_nu || country_oceania || country_polynesia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiue.RegisterName(xlanguage.Russian, "Ниуэ")
	dataNiue.RegisterOfficialName(xlanguage.Russian, "Ниуэ")
	dataNiue.RegisterCapital(xlanguage.Russian, "Алофи")
}
