//go:build (lang_ja || lang_all) && (country_all || country_americas || country_ca || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCanada.RegisterName(xlanguage.Japanese, "カナダ")
	dataCanada.RegisterOfficialName(xlanguage.Japanese, "カナダ")
	dataCanada.RegisterCapital(xlanguage.Japanese, "オタワ")
}
