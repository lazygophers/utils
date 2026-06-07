//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTaiwan.RegisterName(xlanguage.Arabic, "تايوان")
	dataTaiwan.RegisterOfficialName(xlanguage.Arabic, "جمهورية الصين")
	dataTaiwan.RegisterCapital(xlanguage.Arabic, "تايبيه")
}
