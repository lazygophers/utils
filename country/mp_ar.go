//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthernMarianaIslands.RegisterName(xlanguage.Arabic, "جزر ماريانا الشمالية")
	dataNorthernMarianaIslands.RegisterOfficialName(xlanguage.Arabic, "كومنولث جزر ماريانا الشمالية")
	dataNorthernMarianaIslands.RegisterCapital(xlanguage.Arabic, "سايبان")
}
