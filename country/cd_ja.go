//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDrCongo.RegisterName(xlanguage.Japanese, "コンゴ民主共和国")
	dataDrCongo.RegisterOfficialName(xlanguage.Japanese, "コンゴ民主共和国")
	dataDrCongo.RegisterCapital(xlanguage.Japanese, "キンシャサ")
}
