//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsVirginIslands.RegisterName(xlanguage.Arabic, "جزر العذراء الأمريكية")
	dataUsVirginIslands.RegisterOfficialName(xlanguage.Arabic, "إقليم جزر العذراء الأمريكية")
	dataUsVirginIslands.RegisterCapital(xlanguage.Arabic, "شارلوت أمالي")
}
