//go:build country_all || country_europe || country_li || country_western_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiechtenstein.RegisterName(xlanguage.Chinese, "列支敦士登")
	dataLiechtenstein.RegisterOfficialName(xlanguage.Chinese, "列支敦士登公国")
	dataLiechtenstein.RegisterCapital(xlanguage.Chinese, "瓦杜兹")
}
