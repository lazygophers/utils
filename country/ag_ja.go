//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntiguaAndBarbuda.RegisterName(xlanguage.Japanese, "アンティグア・バーブーダ")
	dataAntiguaAndBarbuda.RegisterOfficialName(xlanguage.Japanese, "アンティグア・バーブーダ")
	dataAntiguaAndBarbuda.RegisterCapital(xlanguage.Japanese, "セントジョンズ")
}
