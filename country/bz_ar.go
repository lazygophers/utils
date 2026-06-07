//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelize.RegisterName(xlanguage.Arabic, "بليز")
	dataBelize.RegisterOfficialName(xlanguage.Arabic, "بليز")
	dataBelize.RegisterCapital(xlanguage.Arabic, "بلموبان")
}
