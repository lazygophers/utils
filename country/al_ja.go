//go:build (lang_ja || lang_all) && (country_al || country_all || country_europe || country_southern_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlbania.RegisterName(xlanguage.Japanese, "アルバニア")
	dataAlbania.RegisterOfficialName(xlanguage.Japanese, "アルバニア共和国")
	dataAlbania.RegisterCapital(xlanguage.Japanese, "ティラナ")
}
