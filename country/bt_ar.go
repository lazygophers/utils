//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBhutan.RegisterName(xlanguage.Arabic, "بوتان")
	dataBhutan.RegisterOfficialName(xlanguage.Arabic, "مملكة بوتان")
	dataBhutan.RegisterCapital(xlanguage.Arabic, "تيمفو")
}
