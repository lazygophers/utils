//go:build (lang_ja || lang_all) && (country_all || country_antarctic || country_aq)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntarctica.RegisterName(xlanguage.Japanese, "南極大陸")
	dataAntarctica.RegisterOfficialName(xlanguage.Japanese, "南極大陸")
}
