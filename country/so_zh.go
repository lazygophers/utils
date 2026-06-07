//go:build country_africa || country_all || country_eastern_africa || country_so

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSomalia.RegisterName(xlanguage.Chinese, "索马里")
	dataSomalia.RegisterOfficialName(xlanguage.Chinese, "索马里联邦共和国")
	dataSomalia.RegisterCapital(xlanguage.Chinese, "摩加迪沙")
}
