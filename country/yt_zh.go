//go:build country_africa || country_all || country_eastern_africa || country_yt

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMayotte.RegisterName(xlanguage.Chinese, "马约特")
	dataMayotte.RegisterOfficialName(xlanguage.Chinese, "马约特")
	dataMayotte.RegisterCapital(xlanguage.Chinese, "马穆楚")
}
