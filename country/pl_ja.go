//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPoland.RegisterName(xlanguage.Japanese, "ポーランド")
	dataPoland.RegisterOfficialName(xlanguage.Japanese, "ポーランド共和国")
	dataPoland.RegisterCapital(xlanguage.Japanese, "ワルシャワ")
}
