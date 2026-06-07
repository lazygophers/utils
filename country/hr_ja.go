//go:build (lang_ja || lang_all) && (country_all || country_europe || country_hr || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCroatia.RegisterName(xlanguage.Japanese, "クロアチア")
	dataCroatia.RegisterOfficialName(xlanguage.Japanese, "クロアチア共和国")
	dataCroatia.RegisterCapital(xlanguage.Japanese, "ザグレブ")
}
