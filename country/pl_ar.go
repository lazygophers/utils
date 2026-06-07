//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPoland.RegisterName(xlanguage.Arabic, "بولندا")
	dataPoland.RegisterOfficialName(xlanguage.Arabic, "جمهورية بولندا")
	dataPoland.RegisterCapital(xlanguage.Arabic, "وارسو")
}
