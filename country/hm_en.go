//go:build country_all || country_antarctic || country_hm

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHeardAndMcDonaldIslands.RegisterName(xlanguage.English, "Heard Island and McDonald Islands")
	dataHeardAndMcDonaldIslands.RegisterOfficialName(xlanguage.English, "Territory of Heard Island and McDonald Islands")
}
