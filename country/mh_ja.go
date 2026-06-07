//go:build (lang_ja || lang_all) && (country_all || country_mh || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMarshallIslands.RegisterName(xlanguage.Japanese, "マーシャル諸島")
	dataMarshallIslands.RegisterOfficialName(xlanguage.Japanese, "マーシャル諸島共和国")
	dataMarshallIslands.RegisterCapital(xlanguage.Japanese, "マジュロ")
}
