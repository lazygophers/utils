//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUkraine.RegisterName(xlanguage.Japanese, "ウクライナ")
	dataUkraine.RegisterOfficialName(xlanguage.Japanese, "ウクライナ")
	dataUkraine.RegisterCapital(xlanguage.Japanese, "キーウ")
}
