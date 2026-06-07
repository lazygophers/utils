//go:build (lang_ja || lang_all) && (country_africa || country_all || country_northern_africa || country_tn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTunisia.RegisterName(xlanguage.Japanese, "チュニジア")
	dataTunisia.RegisterOfficialName(xlanguage.Japanese, "チュニジア共和国")
	dataTunisia.RegisterCapital(xlanguage.Japanese, "チュニス")
}
