//go:build (lang_ja || lang_all) && (country_ae || country_all || country_asia || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataUnitedArabEmirates.RegisterName(xlanguage.Japanese, "アラブ首長国連邦")
	dataUnitedArabEmirates.RegisterOfficialName(xlanguage.Japanese, "アラブ首長国連邦")
	dataUnitedArabEmirates.RegisterCapital(xlanguage.Japanese, "アブダビ")
}
