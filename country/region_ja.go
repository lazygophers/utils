//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	RegionEasternAsia.RegisterName(xlanguage.Japanese, "東アジア")
	RegionSouthEasternAsia.RegisterName(xlanguage.Japanese, "東南アジア")
	RegionSouthernAsia.RegisterName(xlanguage.Japanese, "南アジア")
	RegionWesternAsia.RegisterName(xlanguage.Japanese, "西アジア")
	RegionCentralAsia.RegisterName(xlanguage.Japanese, "中央アジア")
	RegionEasternEurope.RegisterName(xlanguage.Japanese, "東ヨーロッパ")
	RegionNorthernEurope.RegisterName(xlanguage.Japanese, "北ヨーロッパ")
	RegionSouthernEurope.RegisterName(xlanguage.Japanese, "南ヨーロッパ")
	RegionWesternEurope.RegisterName(xlanguage.Japanese, "西ヨーロッパ")
	RegionNorthernAfrica.RegisterName(xlanguage.Japanese, "北アフリカ")
	RegionEasternAfrica.RegisterName(xlanguage.Japanese, "東アフリカ")
	RegionMiddleAfrica.RegisterName(xlanguage.Japanese, "中部アフリカ")
	RegionSouthernAfrica.RegisterName(xlanguage.Japanese, "南部アフリカ")
	RegionWesternAfrica.RegisterName(xlanguage.Japanese, "西アフリカ")
	RegionNorthernAmerica.RegisterName(xlanguage.Japanese, "北アメリカ")
	RegionCentralAmerica.RegisterName(xlanguage.Japanese, "中央アメリカ")
	RegionSouthAmerica.RegisterName(xlanguage.Japanese, "南アメリカ")
	RegionCaribbean.RegisterName(xlanguage.Japanese, "カリブ海地域")
	RegionAustraliaAndNewZealand.RegisterName(xlanguage.Japanese, "オーストラリアとニュージーランド")
	RegionMelanesia.RegisterName(xlanguage.Japanese, "メラネシア")
	RegionMicronesia.RegisterName(xlanguage.Japanese, "ミクロネシア")
	RegionPolynesia.RegisterName(xlanguage.Japanese, "ポリネシア")
	RegionAntarctic.RegisterName(xlanguage.Japanese, "南極地域")
}
