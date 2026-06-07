//go:build (lang_ja || lang_all) && (country_all || country_ki || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKiribati.RegisterName(xlanguage.Japanese, "キリバス")
	dataKiribati.RegisterOfficialName(xlanguage.Japanese, "キリバス共和国")
	dataKiribati.RegisterCapital(xlanguage.Japanese, "タラワ")
}
