//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_pr)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPuertoRico.RegisterName(xlanguage.Arabic, "بورتوريكو")
	dataPuertoRico.RegisterOfficialName(xlanguage.Arabic, "كومنولث بورتوريكو")
	dataPuertoRico.RegisterCapital(xlanguage.Arabic, "سان خوان")
}
