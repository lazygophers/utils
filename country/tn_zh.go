//go:build country_africa || country_all || country_northern_africa || country_tn

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTunisia.RegisterName(xlanguage.Chinese, "突尼斯")
	dataTunisia.RegisterOfficialName(xlanguage.Chinese, "突尼斯共和国")
	dataTunisia.RegisterCapital(xlanguage.Chinese, "突尼斯市")
}
