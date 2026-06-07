//go:build (lang_ja || lang_all) && (country_all || country_antarctic || country_tf)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFrenchSouthernTerritories.RegisterName(xlanguage.Japanese, "フランス領南方・南極地域")
	dataFrenchSouthernTerritories.RegisterOfficialName(xlanguage.Japanese, "フランス領南方・南極地域")
	dataFrenchSouthernTerritories.RegisterCapital(xlanguage.Japanese, "サン＝ピエール")
}
