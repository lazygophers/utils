//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCentralAfricanRepublic.RegisterName(xlanguage.Japanese, "中央アフリカ共和国")
	dataCentralAfricanRepublic.RegisterOfficialName(xlanguage.Japanese, "中央アフリカ共和国")
	dataCentralAfricanRepublic.RegisterCapital(xlanguage.Japanese, "バンギ")
}
