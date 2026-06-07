//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataReunion.RegisterName(xlanguage.Arabic, "لا ريونيون")
	dataReunion.RegisterOfficialName(xlanguage.Arabic, "لا ريونيون")
	dataReunion.RegisterCapital(xlanguage.Arabic, "سان دوني")
}
