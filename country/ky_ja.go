//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_ky)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaymanIslands.RegisterName(xlanguage.Japanese, "ケイマン諸島")
	dataCaymanIslands.RegisterOfficialName(xlanguage.Japanese, "ケイマン諸島")
	dataCaymanIslands.RegisterCapital(xlanguage.Japanese, "ジョージタウン")
}
