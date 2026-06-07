//go:build lang_es || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHongKong.RegisterName(xlanguage.Spanish, "Hong Kong")
	dataHongKong.RegisterOfficialName(xlanguage.Spanish, "Región Administrativa Especial de Hong Kong")
	dataHongKong.RegisterCapital(xlanguage.Spanish, "Hong Kong")
}
