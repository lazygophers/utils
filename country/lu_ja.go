//go:build (lang_ja || lang_all) && (country_all || country_europe || country_lu || country_western_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLuxembourg.RegisterName(xlanguage.Japanese, "ルクセンブルク")
	dataLuxembourg.RegisterOfficialName(xlanguage.Japanese, "ルクセンブルク大公国")
	dataLuxembourg.RegisterCapital(xlanguage.Japanese, "ルクセンブルク市")
}
