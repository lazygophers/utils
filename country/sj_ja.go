//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSvalbardAndJanMayen.RegisterName(xlanguage.Japanese, "スヴァールバル諸島およびヤンマイエン島")
	dataSvalbardAndJanMayen.RegisterOfficialName(xlanguage.Japanese, "スヴァールバル諸島およびヤンマイエン島")
	dataSvalbardAndJanMayen.RegisterCapital(xlanguage.Japanese, "ロングイェールビーン")
}
