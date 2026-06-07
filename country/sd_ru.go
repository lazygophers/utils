//go:build (lang_ru || lang_all) && (country_africa || country_all || country_northern_africa || country_sd)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSudan.RegisterName(xlanguage.Russian, "Судан")
	dataSudan.RegisterOfficialName(xlanguage.Russian, "Республика Судан")
	dataSudan.RegisterCapital(xlanguage.Russian, "Хартум")
}
