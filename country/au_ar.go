//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustralia.RegisterName(xlanguage.Arabic, "أستراليا")
	dataAustralia.RegisterOfficialName(xlanguage.Arabic, "كومنولث أستراليا")
	dataAustralia.RegisterCapital(xlanguage.Arabic, "كانبرا")
}
