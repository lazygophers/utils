//go:build country_all || country_fm || country_micronesia || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMicronesia.RegisterName(xlanguage.Chinese, "密克罗尼西亚联邦")
	dataMicronesia.RegisterOfficialName(xlanguage.Chinese, "密克罗尼西亚联邦")
	dataMicronesia.RegisterCapital(xlanguage.Chinese, "帕利基尔")
}
