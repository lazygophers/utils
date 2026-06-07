//go:build (lang_ru || lang_all) && (country_all || country_antarctic || country_hm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHeardAndMcDonaldIslands.RegisterName(xlanguage.Russian, "Остров Херд и острова Макдональд")
	dataHeardAndMcDonaldIslands.RegisterOfficialName(xlanguage.Russian, "Территория Остров Херд и острова Макдональд")
}
