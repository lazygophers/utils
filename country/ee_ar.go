//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEstonia.RegisterName(xlanguage.Arabic, "إستونيا")
	dataEstonia.RegisterOfficialName(xlanguage.Arabic, "جمهورية إستونيا")
	dataEstonia.RegisterCapital(xlanguage.Arabic, "تالين")
}
