//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGibraltar.RegisterName(xlanguage.Japanese, "ジブラルタル")
	dataGibraltar.RegisterOfficialName(xlanguage.Japanese, "ジブラルタル")
	dataGibraltar.RegisterCapital(xlanguage.Japanese, "ジブラルタル")
}
