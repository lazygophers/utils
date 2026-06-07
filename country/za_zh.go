//go:build country_africa || country_all || country_southern_africa || country_za

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthAfrica.RegisterName(xlanguage.Chinese, "南非")
	dataSouthAfrica.RegisterOfficialName(xlanguage.Chinese, "南非共和国")
	dataSouthAfrica.RegisterCapital(xlanguage.Chinese, "比勒陀利亚")
}
