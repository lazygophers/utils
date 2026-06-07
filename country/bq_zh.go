//go:build country_all || country_americas || country_bq || country_caribbean

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBonaireSintEustatiusAndSaba.RegisterName(xlanguage.Chinese, "荷兰加勒比区")
	dataBonaireSintEustatiusAndSaba.RegisterOfficialName(xlanguage.Chinese, "博奈尔、圣尤斯特歇斯和萨巴")
}
