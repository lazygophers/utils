//go:build country_all || country_micronesia || country_oceania || country_um

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUsMinorOutlyingIslands.RegisterName(xlanguage.Chinese, "美国本土外小岛屿")
	dataUsMinorOutlyingIslands.RegisterOfficialName(xlanguage.Chinese, "美国本土外小岛屿")
}
