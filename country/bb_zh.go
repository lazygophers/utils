//go:build country_all || country_americas || country_bb || country_caribbean

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBarbados.RegisterName(xlanguage.Chinese, "巴巴多斯")
	dataBarbados.RegisterOfficialName(xlanguage.Chinese, "巴巴多斯")
	dataBarbados.RegisterCapital(xlanguage.Chinese, "布里奇敦")
}
