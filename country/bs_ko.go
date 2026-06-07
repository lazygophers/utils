//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahamas.RegisterName(xlanguage.Korean, "바하마")
	dataBahamas.RegisterOfficialName(xlanguage.Korean, "바하마 연방")
	dataBahamas.RegisterCapital(xlanguage.Korean, "나소")
}
