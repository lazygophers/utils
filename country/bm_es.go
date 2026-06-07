//go:build (lang_es || lang_all) && (country_all || country_americas || country_bm || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBermuda.RegisterName(xlanguage.Spanish, "Bermudas")
	dataBermuda.RegisterOfficialName(xlanguage.Spanish, "Bermudas")
	dataBermuda.RegisterCapital(xlanguage.Spanish, "Hamilton")
}
