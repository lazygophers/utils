//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPalau.RegisterName(xlanguage.Japanese, "パラオ")
	dataPalau.RegisterOfficialName(xlanguage.Japanese, "パラオ共和国")
	dataPalau.RegisterCapital(xlanguage.Japanese, "マルキョク")
}
