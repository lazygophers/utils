//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishVirginIslands.RegisterName(xlanguage.Japanese, "イギリス領ヴァージン諸島")
	dataBritishVirginIslands.RegisterOfficialName(xlanguage.Japanese, "ヴァージン諸島")
	dataBritishVirginIslands.RegisterCapital(xlanguage.Japanese, "ロードタウン")
}
