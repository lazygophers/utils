//go:build country_africa || country_all || country_eastern_africa || country_yt

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMayotte.RegisterName(xlanguage.French, "Mayotte")
	dataMayotte.RegisterOfficialName(xlanguage.French, "Département de Mayotte")
	dataMayotte.RegisterCapital(xlanguage.French, "Mamoudzou")
}
