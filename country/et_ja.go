//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_et)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEthiopia.RegisterName(xlanguage.Japanese, "エチオピア")
	dataEthiopia.RegisterOfficialName(xlanguage.Japanese, "エチオピア連邦民主共和国")
	dataEthiopia.RegisterCapital(xlanguage.Japanese, "アディスアベバ")
}
