//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSingapore.RegisterName(xlanguage.Arabic, "سنغافورة")
	dataSingapore.RegisterOfficialName(xlanguage.Arabic, "جمهورية سنغافورة")
	dataSingapore.RegisterCapital(xlanguage.Arabic, "سنغافورة")
}
