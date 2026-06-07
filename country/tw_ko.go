//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTaiwan.RegisterName(xlanguage.Korean, "대만")
	dataTaiwan.RegisterOfficialName(xlanguage.Korean, "중화민국")
	dataTaiwan.RegisterCapital(xlanguage.Korean, "타이베이")
}
