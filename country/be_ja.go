//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelgium.RegisterName(xlanguage.Japanese, "ベルギー")
	dataBelgium.RegisterOfficialName(xlanguage.Japanese, "ベルギー王国")
	dataBelgium.RegisterCapital(xlanguage.Japanese, "ブリュッセル")
}
