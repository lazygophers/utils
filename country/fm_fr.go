//go:build (lang_fr || lang_all) && (country_all || country_fm || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMicronesia.RegisterName(xlanguage.French, "États fédérés de Micronésie")
	dataMicronesia.RegisterOfficialName(xlanguage.French, "États fédérés de Micronésie")
	dataMicronesia.RegisterCapital(xlanguage.French, "Palikir")
}
