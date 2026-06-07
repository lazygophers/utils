//go:build (lang_ru || lang_all) && (country_all || country_oceania || country_polynesia || country_tk)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTokelau.RegisterName(xlanguage.Russian, "Токелау")
	dataTokelau.RegisterOfficialName(xlanguage.Russian, "Токелау")
}
