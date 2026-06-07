//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSierraLeone.RegisterName(xlanguage.Arabic, "سيراليون")
	dataSierraLeone.RegisterOfficialName(xlanguage.Arabic, "جمهورية سيراليون")
	dataSierraLeone.RegisterCapital(xlanguage.Arabic, "فريتاون")
}
