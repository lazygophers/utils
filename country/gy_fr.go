//go:build (lang_fr || lang_all) && (country_all || country_americas || country_gy || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuyana.RegisterName(xlanguage.French, "Guyana")
	dataGuyana.RegisterOfficialName(xlanguage.French, "République coopérative du Guyana")
	dataGuyana.RegisterCapital(xlanguage.French, "Georgetown")
}
