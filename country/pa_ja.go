//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPanama.RegisterName(xlanguage.Japanese, "パナマ")
	dataPanama.RegisterOfficialName(xlanguage.Japanese, "パナマ共和国")
	dataPanama.RegisterCapital(xlanguage.Japanese, "パナマシティ")
}
