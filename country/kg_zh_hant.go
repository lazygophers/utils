//go:build lang_zh_hant || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKyrgyzstan.RegisterName(xlanguage.MustParse("zh-Hant"), "吉爾吉斯")
	dataKyrgyzstan.RegisterOfficialName(xlanguage.MustParse("zh-Hant"), "吉爾吉斯共和國")
	dataKyrgyzstan.RegisterCapital(xlanguage.MustParse("zh-Hant"), "比斯凱克")
}
