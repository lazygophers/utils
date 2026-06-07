//go:build (lang_ko || lang_all) && (country_all || country_americas || country_bq || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBonaireSintEustatiusAndSaba.RegisterName(xlanguage.Korean, "카리브 네덜란드")
	dataBonaireSintEustatiusAndSaba.RegisterOfficialName(xlanguage.Korean, "보네르, 신트외스타티위스, 사바")
	dataBonaireSintEustatiusAndSaba.RegisterCapital(xlanguage.Korean, "크랄렌데이크")
}
