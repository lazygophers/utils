//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlbania.RegisterName(xlanguage.Arabic, "ألبانيا")
	dataAlbania.RegisterOfficialName(xlanguage.Arabic, "جمهورية ألبانيا")
	dataAlbania.RegisterCapital(xlanguage.Arabic, "تيرانا")
}
