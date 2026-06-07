//go:build country_all || country_americas || country_central_america || country_gt

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuatemala.RegisterName(xlanguage.Chinese, "危地马拉")
	dataGuatemala.RegisterOfficialName(xlanguage.Chinese, "危地马拉共和国")
	dataGuatemala.RegisterCapital(xlanguage.Chinese, "危地马拉城")
}
