//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataRomania.RegisterName(xlanguage.Japanese, "ルーマニア")
	dataRomania.RegisterOfficialName(xlanguage.Japanese, "ルーマニア")
	dataRomania.RegisterCapital(xlanguage.Japanese, "ブカレスト")
}
