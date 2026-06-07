//go:build (lang_ar || lang_all) && (country_all || country_americas || country_bl || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintBarthelemy.RegisterName(xlanguage.Arabic, "سان بارتيلمي")
	dataSaintBarthelemy.RegisterOfficialName(xlanguage.Arabic, "جماعة سان بارتيلمي")
	dataSaintBarthelemy.RegisterCapital(xlanguage.Arabic, "غوستافيا")
}
