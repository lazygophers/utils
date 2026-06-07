//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedArabEmirates.RegisterName(xlanguage.Japanese, "アラブ首長国連邦")
	dataUnitedArabEmirates.RegisterOfficialName(xlanguage.Japanese, "アラブ首長国連邦")
	dataUnitedArabEmirates.RegisterCapital(xlanguage.Japanese, "アブダビ")
}
