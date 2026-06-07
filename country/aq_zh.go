//go:build country_all || country_antarctic || country_aq

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntarctica.RegisterName(xlanguage.Chinese, "南极洲")
	dataAntarctica.RegisterOfficialName(xlanguage.Chinese, "南极洲")
}
