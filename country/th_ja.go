//go:build (lang_ja || lang_all) && (country_all || country_asia || country_south_eastern_asia || country_th)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataThailand.RegisterName(xlanguage.Japanese, "タイ王国")
	dataThailand.RegisterOfficialName(xlanguage.Japanese, "タイ王国")
	dataThailand.RegisterCapital(xlanguage.Japanese, "バンコク")
}
