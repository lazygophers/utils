//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_mf)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintMartin.RegisterName(xlanguage.Arabic, "سانت مارتن")
	dataSaintMartin.RegisterOfficialName(xlanguage.Arabic, "جماعة سانت مارتن")
	dataSaintMartin.RegisterCapital(xlanguage.Arabic, "مارينو")
}
