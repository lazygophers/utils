//go:build (lang_es || lang_all) && (country_all || country_antarctic || country_bv)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBouvetIsland.RegisterName(xlanguage.Spanish, "Isla Bouvet")
	dataBouvetIsland.RegisterOfficialName(xlanguage.Spanish, "Isla Bouvet")
}
