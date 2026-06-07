//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMarshallIslands.RegisterName(xlanguage.Japanese, "マーシャル諸島")
	dataMarshallIslands.RegisterOfficialName(xlanguage.Japanese, "マーシャル諸島共和国")
	dataMarshallIslands.RegisterCapital(xlanguage.Japanese, "マジュロ")
}
