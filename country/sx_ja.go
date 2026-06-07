//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_sx)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSintMaarten.RegisterName(xlanguage.Japanese, "シント・マールテン")
	dataSintMaarten.RegisterOfficialName(xlanguage.Japanese, "シント・マールテン")
	dataSintMaarten.RegisterCapital(xlanguage.Japanese, "フィリップスブルフ")
}
