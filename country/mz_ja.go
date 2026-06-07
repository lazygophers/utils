//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.Japanese, "モザンビーク")
	dataMozambique.RegisterOfficialName(xlanguage.Japanese, "モザンビーク共和国")
	dataMozambique.RegisterCapital(xlanguage.Japanese, "マプト")
}
