//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLuxembourg.RegisterName(xlanguage.Japanese, "ルクセンブルク")
	dataLuxembourg.RegisterOfficialName(xlanguage.Japanese, "ルクセンブルク大公国")
	dataLuxembourg.RegisterCapital(xlanguage.Japanese, "ルクセンブルク市")
}
