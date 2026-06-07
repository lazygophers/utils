//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGhana.RegisterName(xlanguage.Japanese, "ガーナ")
	dataGhana.RegisterOfficialName(xlanguage.Japanese, "ガーナ共和国")
	dataGhana.RegisterCapital(xlanguage.Japanese, "アクラ")
}
