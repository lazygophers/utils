//go:build country_all || country_asia || country_eastern_asia || country_mn || currency_all || currency_mnt

package currency

import xlanguage "golang.org/x/text/language"

func init() {
	MNT.RegisterName(xlanguage.Chinese, "蒙古图格里克")
}
