//go:build (lang_ar || lang_all) && (country_all || country_ax || country_europe || country_northern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlandIslands.RegisterName(xlanguage.Arabic, "جزر أولاند")
	dataAlandIslands.RegisterOfficialName(xlanguage.Arabic, "جزر أولاند")
	dataAlandIslands.RegisterCapital(xlanguage.Arabic, "ماريهامن")
}
