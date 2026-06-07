//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTokelau.RegisterName(xlanguage.Arabic, "توكيلاو")
	dataTokelau.RegisterOfficialName(xlanguage.Arabic, "توكيلاو")
	dataTokelau.RegisterCapital(xlanguage.Arabic, "—")
}
