//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_vg)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishVirginIslands.RegisterName(xlanguage.Japanese, "イギリス領ヴァージン諸島")
	dataBritishVirginIslands.RegisterOfficialName(xlanguage.Japanese, "ヴァージン諸島")
	dataBritishVirginIslands.RegisterCapital(xlanguage.Japanese, "ロードタウン")
}
