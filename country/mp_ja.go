//go:build (lang_ja || lang_all) && (country_all || country_micronesia || country_mp || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNorthernMarianaIslands.RegisterName(xlanguage.Japanese, "北マリアナ諸島")
	dataNorthernMarianaIslands.RegisterOfficialName(xlanguage.Japanese, "北マリアナ諸島")
	dataNorthernMarianaIslands.RegisterCapital(xlanguage.Japanese, "サイパン")
}
