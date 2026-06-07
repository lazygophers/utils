//go:build country_all || country_americas || country_bz || country_central_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelize.RegisterName(xlanguage.Chinese, "伯利兹")
	dataBelize.RegisterOfficialName(xlanguage.Chinese, "伯利兹")
	dataBelize.RegisterCapital(xlanguage.Chinese, "贝尔莫潘")
}
