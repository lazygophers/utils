//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_mu)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritius.RegisterName(xlanguage.Japanese, "モーリシャス")
	dataMauritius.RegisterOfficialName(xlanguage.Japanese, "モーリシャス共和国")
	dataMauritius.RegisterCapital(xlanguage.Japanese, "ポートルイス")
}
