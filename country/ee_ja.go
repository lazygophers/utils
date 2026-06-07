//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataEstonia.RegisterName(xlanguage.Japanese, "エストニア")
	dataEstonia.RegisterOfficialName(xlanguage.Japanese, "エストニア共和国")
	dataEstonia.RegisterCapital(xlanguage.Japanese, "タリン")
}
