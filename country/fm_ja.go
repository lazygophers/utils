//go:build (lang_ja || lang_all) && (country_all || country_fm || country_micronesia || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMicronesia.RegisterName(xlanguage.Japanese, "ミクロネシア連邦")
	dataMicronesia.RegisterOfficialName(xlanguage.Japanese, "ミクロネシア連邦")
	dataMicronesia.RegisterCapital(xlanguage.Japanese, "パリキール")
}
