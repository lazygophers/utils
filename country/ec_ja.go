//go:build (lang_ja || lang_all) && (country_all || country_americas || country_ec || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEcuador.RegisterName(xlanguage.Japanese, "エクアドル")
	dataEcuador.RegisterOfficialName(xlanguage.Japanese, "エクアドル共和国")
	dataEcuador.RegisterCapital(xlanguage.Japanese, "キト")
}
