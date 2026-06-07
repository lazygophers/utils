//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHongKong.RegisterName(xlanguage.French, "Hong Kong")
	dataHongKong.RegisterOfficialName(xlanguage.French, "Région administrative spéciale de Hong Kong")
	dataHongKong.RegisterCapital(xlanguage.French, "Hong Kong")
}
