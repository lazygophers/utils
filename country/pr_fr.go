//go:build (lang_fr || lang_all) && (country_all || country_americas || country_caribbean || country_pr)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPuertoRico.RegisterName(xlanguage.French, "Porto Rico")
	dataPuertoRico.RegisterOfficialName(xlanguage.French, "Commonwealth de Porto Rico")
	dataPuertoRico.RegisterCapital(xlanguage.French, "San Juan")
}
