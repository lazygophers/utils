//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCroatia.RegisterName(xlanguage.Japanese, "クロアチア")
	dataCroatia.RegisterOfficialName(xlanguage.Japanese, "クロアチア共和国")
	dataCroatia.RegisterCapital(xlanguage.Japanese, "ザグレブ")
}
