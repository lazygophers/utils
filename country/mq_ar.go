//go:build (lang_ar || lang_all) && (country_all || country_americas || country_caribbean || country_mq)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMartinique.RegisterName(xlanguage.Arabic, "مارتينيك")
	dataMartinique.RegisterOfficialName(xlanguage.Arabic, "مارتينيك")
	dataMartinique.RegisterCapital(xlanguage.Arabic, "فور دو فرانس")
}
