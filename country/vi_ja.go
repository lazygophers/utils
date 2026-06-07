//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsVirginIslands.RegisterName(xlanguage.Japanese, "アメリカ領ヴァージン諸島")
	dataUsVirginIslands.RegisterOfficialName(xlanguage.Japanese, "アメリカ領ヴァージン諸島")
	dataUsVirginIslands.RegisterCapital(xlanguage.Japanese, "シャーロット・アマリー")
}
