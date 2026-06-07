//go:build (lang_fr || lang_all) && (country_all || country_americas || country_bm || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBermuda.RegisterName(xlanguage.French, "Bermudes")
	dataBermuda.RegisterOfficialName(xlanguage.French, "Bermudes")
	dataBermuda.RegisterCapital(xlanguage.French, "Hamilton")
}
