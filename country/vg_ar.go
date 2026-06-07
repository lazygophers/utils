//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_vg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishVirginIslands.RegisterName(xlanguage.Arabic, "جزر العذراء البريطانية")
	dataBritishVirginIslands.RegisterOfficialName(xlanguage.Arabic, "إقليم جزر العذراء البريطاني فيما وراء البحار")
	dataBritishVirginIslands.RegisterCapital(xlanguage.Arabic, "رود تاون")
}
