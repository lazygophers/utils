//go:build country_africa || country_all || country_eastern_africa || country_mz

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMozambique.RegisterName(xlanguage.Chinese, "莫桑比克")
	dataMozambique.RegisterOfficialName(xlanguage.Chinese, "莫桑比克共和国")
	dataMozambique.RegisterCapital(xlanguage.Chinese, "马普托")
}
