//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalau.RegisterName(xlanguage.Korean, "팔라우")
	dataPalau.RegisterOfficialName(xlanguage.Korean, "팔라우 공화국")
	dataPalau.RegisterCapital(xlanguage.Korean, "응게룰무드")
}
