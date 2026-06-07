//go:build country_all || country_eastern_europe || country_europe || country_md

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMoldova.RegisterName(xlanguage.Chinese, "摩尔多瓦")
	dataMoldova.RegisterOfficialName(xlanguage.Chinese, "摩尔多瓦共和国")
	dataMoldova.RegisterCapital(xlanguage.Chinese, "基希讷乌")
}
