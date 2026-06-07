//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalau.RegisterName(xlanguage.Arabic, "بالاو")
	dataPalau.RegisterOfficialName(xlanguage.Arabic, "جمهورية بالاو")
	dataPalau.RegisterCapital(xlanguage.Arabic, "نغرولمود")
}
