//go:build country_all || country_oceania || country_polynesia || country_to

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataTonga.RegisterName(xlanguage.Chinese, "汤加")
	dataTonga.RegisterOfficialName(xlanguage.Chinese, "汤加王国")
	dataTonga.RegisterCapital(xlanguage.Chinese, "努库阿洛法")
}
