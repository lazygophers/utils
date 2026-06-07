//go:build (lang_ja || lang_all) && (country_all || country_europe || country_northern_europe || country_sj)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSvalbardAndJanMayen.RegisterName(xlanguage.Japanese, "スヴァールバル諸島およびヤンマイエン島")
	dataSvalbardAndJanMayen.RegisterOfficialName(xlanguage.Japanese, "スヴァールバル諸島およびヤンマイエン島")
	dataSvalbardAndJanMayen.RegisterCapital(xlanguage.Japanese, "ロングイェールビーン")
}
