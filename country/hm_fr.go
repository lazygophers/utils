//go:build lang_fr || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHeardAndMcDonaldIslands.RegisterName(xlanguage.French, "Îles Heard-et-MacDonald")
	dataHeardAndMcDonaldIslands.RegisterOfficialName(xlanguage.French, "Territoire des îles Heard-et-MacDonald")
}
