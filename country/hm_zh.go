//go:build country_all || country_antarctic || country_hm

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHeardAndMcDonaldIslands.RegisterName(xlanguage.Chinese, "赫德岛和麦克唐纳群岛")
	dataHeardAndMcDonaldIslands.RegisterOfficialName(xlanguage.Chinese, "赫德岛和麦克唐纳群岛领地")
}
