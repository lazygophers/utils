//go:build (lang_ar || lang_all) && (country_all || country_ee || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEstonia.RegisterName(xlanguage.Arabic, "إستونيا")
	dataEstonia.RegisterOfficialName(xlanguage.Arabic, "جمهورية إستونيا")
	dataEstonia.RegisterCapital(xlanguage.Arabic, "تالين")
}
