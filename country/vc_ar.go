//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintVincentAndGrenadines.RegisterName(xlanguage.Arabic, "سانت فنسنت والغرينادين")
	dataSaintVincentAndGrenadines.RegisterOfficialName(xlanguage.Arabic, "سانت فنسنت والغرينادين")
	dataSaintVincentAndGrenadines.RegisterCapital(xlanguage.Arabic, "كينغستاون")
}
