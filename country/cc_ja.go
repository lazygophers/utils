//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCocosKeelingIslands.RegisterName(xlanguage.Japanese, "ココス諸島")
	dataCocosKeelingIslands.RegisterOfficialName(xlanguage.Japanese, "ココス（キーリング）諸島")
	dataCocosKeelingIslands.RegisterCapital(xlanguage.Japanese, "ウェスト島")
}
