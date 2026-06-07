//go:build country_all || country_antarctic || country_bv

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBouvetIsland.RegisterName(xlanguage.English, "Bouvet Island")
	dataBouvetIsland.RegisterOfficialName(xlanguage.English, "Bouvet Island")
}
