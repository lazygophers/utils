//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVenezuela.RegisterName(xlanguage.Japanese, "ベネズエラ")
	dataVenezuela.RegisterOfficialName(xlanguage.Japanese, "ベネズエラ・ボリバル共和国")
	dataVenezuela.RegisterCapital(xlanguage.Japanese, "カラカス")
}
