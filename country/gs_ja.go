//go:build (lang_ja || lang_all) && (country_all || country_antarctic || country_gs)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthGeorgiaAndSouthSandwich.RegisterName(xlanguage.Japanese, "サウスジョージア・サウスサンドウィッチ諸島")
	dataSouthGeorgiaAndSouthSandwich.RegisterOfficialName(xlanguage.Japanese, "サウスジョージア・サウスサンドウィッチ諸島")
	dataSouthGeorgiaAndSouthSandwich.RegisterCapital(xlanguage.Japanese, "キング・エドワード・ポイント")
}
