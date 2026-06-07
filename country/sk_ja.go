//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSlovakia.RegisterName(xlanguage.Japanese, "スロバキア")
	dataSlovakia.RegisterOfficialName(xlanguage.Japanese, "スロバキア共和国")
	dataSlovakia.RegisterCapital(xlanguage.Japanese, "ブラチスラヴァ")
}
