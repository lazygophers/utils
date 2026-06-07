//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_dm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominica.RegisterName(xlanguage.Arabic, "دومينيكا")
	dataDominica.RegisterOfficialName(xlanguage.Arabic, "كومنولث دومينيكا")
	dataDominica.RegisterCapital(xlanguage.Arabic, "روسو")
}
