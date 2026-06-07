//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTajikistan.RegisterName(xlanguage.Japanese, "タジキスタン")
	dataTajikistan.RegisterOfficialName(xlanguage.Japanese, "タジキスタン共和国")
	dataTajikistan.RegisterCapital(xlanguage.Japanese, "ドゥシャンベ")
}
