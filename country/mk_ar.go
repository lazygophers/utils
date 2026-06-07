//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthMacedonia.RegisterName(xlanguage.Arabic, "مقدونيا الشمالية")
	dataNorthMacedonia.RegisterOfficialName(xlanguage.Arabic, "جمهورية مقدونيا الشمالية")
	dataNorthMacedonia.RegisterCapital(xlanguage.Arabic, "سكوبيه")
}
