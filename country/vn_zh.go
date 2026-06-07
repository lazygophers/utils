//go:build country_all || country_asia || country_south_eastern_asia || country_vn

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataVietnam.RegisterName(xlanguage.Chinese, "越南")
	dataVietnam.RegisterOfficialName(xlanguage.Chinese, "越南社会主义共和国")
	dataVietnam.RegisterCapital(xlanguage.Chinese, "河内")
}
