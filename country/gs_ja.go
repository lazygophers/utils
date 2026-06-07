//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthGeorgiaAndSouthSandwich.RegisterName(xlanguage.Japanese, "サウスジョージア・サウスサンドウィッチ諸島")
	dataSouthGeorgiaAndSouthSandwich.RegisterOfficialName(xlanguage.Japanese, "サウスジョージア・サウスサンドウィッチ諸島")
	dataSouthGeorgiaAndSouthSandwich.RegisterCapital(xlanguage.Japanese, "キング・エドワード・ポイント")
}
