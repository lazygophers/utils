//go:build (lang_fr || lang_all) && (country_all || country_australia_and_new_zealand || country_cc || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCocosKeelingIslands.RegisterName(xlanguage.French, "Îles Cocos")
	dataCocosKeelingIslands.RegisterOfficialName(xlanguage.French, "Territoire des îles Cocos")
	dataCocosKeelingIslands.RegisterCapital(xlanguage.French, "West Island")
}
