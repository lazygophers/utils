//go:build (lang_fr || lang_all) && (country_all || country_antarctic || country_bv)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBouvetIsland.RegisterName(xlanguage.French, "Île Bouvet")
	dataBouvetIsland.RegisterOfficialName(xlanguage.French, "Île Bouvet")
}
