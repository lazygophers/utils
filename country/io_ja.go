//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBritishIndianOceanTerritory.RegisterName(xlanguage.Japanese, "イギリス領インド洋地域")
	dataBritishIndianOceanTerritory.RegisterOfficialName(xlanguage.Japanese, "イギリス領インド洋地域")
	dataBritishIndianOceanTerritory.RegisterCapital(xlanguage.Japanese, "ディエゴガルシア島")
}
