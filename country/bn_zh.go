//go:build country_all || country_asia || country_bn || country_south_eastern_asia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrunei.RegisterName(xlanguage.Chinese, "文莱")
	dataBrunei.RegisterOfficialName(xlanguage.Chinese, "文莱达鲁萨兰国")
	dataBrunei.RegisterCapital(xlanguage.Chinese, "斯里巴加湾市")
}
