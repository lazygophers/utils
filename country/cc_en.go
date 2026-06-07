//go:build country_all || country_australia_and_new_zealand || country_cc || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCocosKeelingIslands.RegisterName(xlanguage.English, "Cocos (Keeling) Islands")
	dataCocosKeelingIslands.RegisterOfficialName(xlanguage.English, "Territory of the Cocos (Keeling) Islands")
	dataCocosKeelingIslands.RegisterCapital(xlanguage.English, "West Island")
}
