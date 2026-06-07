//go:build country_all || country_antarctic || country_bv

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBouvetIsland.RegisterName(xlanguage.Chinese, "布维岛")
	dataBouvetIsland.RegisterOfficialName(xlanguage.Chinese, "布维岛")
}
