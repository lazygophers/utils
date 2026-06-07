//go:build country_all || country_fj || country_melanesia || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataFiji.RegisterName(xlanguage.Chinese, "斐济")
	dataFiji.RegisterOfficialName(xlanguage.Chinese, "斐济共和国")
	dataFiji.RegisterCapital(xlanguage.Chinese, "苏瓦")
}
