//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMyanmar.RegisterName(xlanguage.Japanese, "ミャンマー")
	dataMyanmar.RegisterOfficialName(xlanguage.Japanese, "ミャンマー連邦共和国")
	dataMyanmar.RegisterCapital(xlanguage.Japanese, "ネピドー")
}
