//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_so)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSomalia.RegisterName(xlanguage.Japanese, "ソマリア")
	dataSomalia.RegisterOfficialName(xlanguage.Japanese, "ソマリア連邦共和国")
	dataSomalia.RegisterCapital(xlanguage.Japanese, "モガディシュ")
}
