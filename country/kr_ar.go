//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthKorea.RegisterName(xlanguage.Arabic, "كوريا الجنوبية")
	dataSouthKorea.RegisterOfficialName(xlanguage.Arabic, "جمهورية كوريا")
	dataSouthKorea.RegisterCapital(xlanguage.Arabic, "سول")
}
