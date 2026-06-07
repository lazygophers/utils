//go:build (lang_ru || lang_all) && (country_all || country_micronesia || country_oceania || country_pw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalau.RegisterName(xlanguage.Russian, "Палау")
	dataPalau.RegisterOfficialName(xlanguage.Russian, "Республика Палау")
	dataPalau.RegisterCapital(xlanguage.Russian, "Нгерулмуд")
}
