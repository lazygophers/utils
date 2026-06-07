//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMontenegro.RegisterName(xlanguage.Japanese, "モンテネグロ")
	dataMontenegro.RegisterOfficialName(xlanguage.Japanese, "モンテネグロ")
	dataMontenegro.RegisterCapital(xlanguage.Japanese, "ポドゴリツァ")
}
