//go:build lang_ko || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMayotte.RegisterName(xlanguage.Korean, "마요트")
	dataMayotte.RegisterOfficialName(xlanguage.Korean, "마요트")
	dataMayotte.RegisterCapital(xlanguage.Korean, "마무주")
}
