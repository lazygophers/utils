//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedKingdom.RegisterName(xlanguage.Japanese, "イギリス")
	dataUnitedKingdom.RegisterOfficialName(xlanguage.Japanese, "グレートブリテン及び北アイルランド連合王国")
	dataUnitedKingdom.RegisterCapital(xlanguage.Japanese, "ロンドン")
}
