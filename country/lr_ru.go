//go:build (lang_ru || lang_all) && (country_africa || country_all || country_lr || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiberia.RegisterName(xlanguage.Russian, "Либерия")
	dataLiberia.RegisterOfficialName(xlanguage.Russian, "Республика Либерия")
	dataLiberia.RegisterCapital(xlanguage.Russian, "Монровия")
}
