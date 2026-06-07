//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataColombia.RegisterName(xlanguage.Japanese, "コロンビア")
	dataColombia.RegisterOfficialName(xlanguage.Japanese, "コロンビア共和国")
	dataColombia.RegisterCapital(xlanguage.Japanese, "ボゴタ")
}
