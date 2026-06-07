//go:build (lang_zh_hant || lang_all) && (country_all || country_americas || country_bq || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBonaireSintEustatiusAndSaba.RegisterName(xlanguage.MustParse("zh-Hant"), "荷蘭加勒比區")
	dataBonaireSintEustatiusAndSaba.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "波奈、聖佑達修斯和薩巴")
	dataBonaireSintEustatiusAndSaba.RegisterCapital(xlanguage.MustParse("zh-Hant"), "克拉倫代克")
}
