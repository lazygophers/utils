//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAruba.RegisterName(xlanguage.Japanese, "アルバ")
	dataAruba.RegisterOfficialName(xlanguage.Japanese, "アルバ")
	dataAruba.RegisterCapital(xlanguage.Japanese, "オラニエスタッド")
}
