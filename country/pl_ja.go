//go:build (lang_ja || lang_all) && (country_all || country_eastern_europe || country_europe || country_pl)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPoland.RegisterName(xlanguage.Japanese, "ポーランド")
	dataPoland.RegisterOfficialName(xlanguage.Japanese, "ポーランド共和国")
	dataPoland.RegisterCapital(xlanguage.Japanese, "ワルシャワ")
}
