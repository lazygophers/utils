//go:build (lang_ru || lang_all) && (country_all || country_antarctic || country_bv)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBouvetIsland.RegisterName(xlanguage.Russian, "Остров Буве")
	dataBouvetIsland.RegisterOfficialName(xlanguage.Russian, "Остров Буве")
}
