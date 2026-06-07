//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGabon.RegisterName(xlanguage.Arabic, "الغابون")
	dataGabon.RegisterOfficialName(xlanguage.Arabic, "الجمهورية الغابونية")
	dataGabon.RegisterCapital(xlanguage.Arabic, "ليبرفيل")
}
