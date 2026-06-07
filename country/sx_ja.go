//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSintMaarten.RegisterName(xlanguage.Japanese, "シント・マールテン")
	dataSintMaarten.RegisterOfficialName(xlanguage.Japanese, "シント・マールテン")
	dataSintMaarten.RegisterCapital(xlanguage.Japanese, "フィリップスブルフ")
}
