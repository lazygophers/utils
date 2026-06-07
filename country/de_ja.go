//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGermany.RegisterName(xlanguage.Japanese, "ドイツ")
	dataGermany.RegisterOfficialName(xlanguage.Japanese, "ドイツ連邦共和国")
	dataGermany.RegisterCapital(xlanguage.Japanese, "ベルリン")
}
