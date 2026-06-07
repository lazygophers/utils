//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.Arabic, "موزمبيق")
	dataMozambique.RegisterOfficialName(xlanguage.Arabic, "جمهورية موزمبيق")
	dataMozambique.RegisterCapital(xlanguage.Arabic, "مابوتو")
}
