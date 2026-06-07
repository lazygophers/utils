//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewCaledonia.RegisterName(xlanguage.Arabic, "كاليدونيا الجديدة")
	dataNewCaledonia.RegisterOfficialName(xlanguage.Arabic, "كاليدونيا الجديدة")
	dataNewCaledonia.RegisterCapital(xlanguage.Arabic, "نوميا")
}
