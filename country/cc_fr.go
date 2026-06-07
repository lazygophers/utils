//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCocosKeelingIslands.RegisterName(xlanguage.French, "Îles Cocos")
	dataCocosKeelingIslands.RegisterOfficialName(xlanguage.French, "Territoire des îles Cocos")
	dataCocosKeelingIslands.RegisterCapital(xlanguage.French, "West Island")
}
