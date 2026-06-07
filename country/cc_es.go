//go:build (lang_es || lang_all) && (country_all || country_australia_and_new_zealand || country_cc || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCocosKeelingIslands.RegisterName(xlanguage.Spanish, "Islas Cocos")
	dataCocosKeelingIslands.RegisterOfficialName(xlanguage.Spanish, "Territorio de las Islas Cocos (Keeling)")
	dataCocosKeelingIslands.RegisterCapital(xlanguage.Spanish, "West Island")
}
