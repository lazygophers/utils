//go:build country_all || country_eastern_europe || country_europe || country_ua

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUkraine.RegisterName(xlanguage.Chinese, "乌克兰")
	dataUkraine.RegisterOfficialName(xlanguage.Chinese, "乌克兰")
	dataUkraine.RegisterCapital(xlanguage.Chinese, "基辅")
}
