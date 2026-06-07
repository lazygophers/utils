//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZambia.RegisterName(xlanguage.Arabic, "زامبيا")
	dataZambia.RegisterOfficialName(xlanguage.Arabic, "جمهورية زامبيا")
	dataZambia.RegisterCapital(xlanguage.Arabic, "لوساكا")
}
