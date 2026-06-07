//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMicronesia.RegisterName(xlanguage.Arabic, "ميكرونيزيا")
	dataMicronesia.RegisterOfficialName(xlanguage.Arabic, "ولايات ميكرونيزيا الموحدة")
	dataMicronesia.RegisterCapital(xlanguage.Arabic, "باليكير")
}
