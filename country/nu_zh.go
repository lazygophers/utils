//go:build country_all || country_nu || country_oceania || country_polynesia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNiue.RegisterName(xlanguage.Chinese, "纽埃")
	dataNiue.RegisterOfficialName(xlanguage.Chinese, "纽埃")
	dataNiue.RegisterCapital(xlanguage.Chinese, "阿洛菲")
}
