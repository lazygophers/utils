//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBotswana.RegisterName(xlanguage.Japanese, "ボツワナ")
	dataBotswana.RegisterOfficialName(xlanguage.Japanese, "ボツワナ共和国")
	dataBotswana.RegisterCapital(xlanguage.Japanese, "ハボローネ")
}
