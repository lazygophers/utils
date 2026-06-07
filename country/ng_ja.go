//go:build (lang_ja || lang_all) && (country_africa || country_all || country_ng || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNigeria.RegisterName(xlanguage.Japanese, "ナイジェリア")
	dataNigeria.RegisterOfficialName(xlanguage.Japanese, "ナイジェリア連邦共和国")
	dataNigeria.RegisterCapital(xlanguage.Japanese, "アブジャ")
}
