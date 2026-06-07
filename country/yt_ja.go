//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_yt)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMayotte.RegisterName(xlanguage.Japanese, "マヨット")
	dataMayotte.RegisterOfficialName(xlanguage.Japanese, "マヨット")
	dataMayotte.RegisterCapital(xlanguage.Japanese, "マムズ")
}
