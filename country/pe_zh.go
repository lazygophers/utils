//go:build country_all || country_americas || country_pe || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPeru.RegisterName(xlanguage.Chinese, "秘鲁")
	dataPeru.RegisterOfficialName(xlanguage.Chinese, "秘鲁共和国")
	dataPeru.RegisterCapital(xlanguage.Chinese, "利马")
}
