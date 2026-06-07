//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCanada.RegisterName(xlanguage.Japanese, "カナダ")
	dataCanada.RegisterOfficialName(xlanguage.Japanese, "カナダ")
	dataCanada.RegisterCapital(xlanguage.Japanese, "オタワ")
}
