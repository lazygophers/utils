//go:build country_all || country_americas || country_gy || country_south_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuyana.RegisterName(xlanguage.Chinese, "圭亚那")
	dataGuyana.RegisterOfficialName(xlanguage.Chinese, "圭亚那合作共和国")
	dataGuyana.RegisterCapital(xlanguage.Chinese, "乔治敦")
}
