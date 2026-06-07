//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBangladesh.RegisterName(xlanguage.Arabic, "بنغلاديش")
	dataBangladesh.RegisterOfficialName(xlanguage.Arabic, "جمهورية بنغلاديش الشعبية")
	dataBangladesh.RegisterCapital(xlanguage.Arabic, "دكا")
}
