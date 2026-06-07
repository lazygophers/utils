//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRwanda.RegisterName(xlanguage.Japanese, "ルワンダ")
	dataRwanda.RegisterOfficialName(xlanguage.Japanese, "ルワンダ共和国")
	dataRwanda.RegisterCapital(xlanguage.Japanese, "キガリ")
}
