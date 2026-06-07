//go:build country_all || country_asia || country_central_asia || country_kg

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataKyrgyzstan.RegisterName(xlanguage.Chinese, "吉尔吉斯斯坦")
	dataKyrgyzstan.RegisterOfficialName(xlanguage.Chinese, "吉尔吉斯共和国")
	dataKyrgyzstan.RegisterCapital(xlanguage.Chinese, "比什凯克")
}
