//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrunei.RegisterName(xlanguage.Arabic, "بروناي")
	dataBrunei.RegisterOfficialName(xlanguage.Arabic, "سلطنة بروناي دار السلام")
	dataBrunei.RegisterCapital(xlanguage.Arabic, "بندر سري بكاوان")
}
