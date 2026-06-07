//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSwitzerland.RegisterName(xlanguage.Japanese, "スイス")
	dataSwitzerland.RegisterOfficialName(xlanguage.Japanese, "スイス連邦")
	dataSwitzerland.RegisterCapital(xlanguage.Japanese, "ベルン")
}
