//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_jm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataJamaica.RegisterName(xlanguage.Japanese, "ジャマイカ")
	dataJamaica.RegisterOfficialName(xlanguage.Japanese, "ジャマイカ")
	dataJamaica.RegisterCapital(xlanguage.Japanese, "キングストン")
}
