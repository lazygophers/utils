//go:build country_all || country_melanesia || country_oceania || country_vu

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVanuatu.RegisterName(xlanguage.Chinese, "瓦努阿图")
	dataVanuatu.RegisterOfficialName(xlanguage.Chinese, "瓦努阿图共和国")
	dataVanuatu.RegisterCapital(xlanguage.Chinese, "维拉港")
}
