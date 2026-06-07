//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintKittsAndNevis.RegisterName(xlanguage.Japanese, "セントクリストファー・ネイビス")
	dataSaintKittsAndNevis.RegisterOfficialName(xlanguage.Japanese, "セントクリストファー・ネイビス")
	dataSaintKittsAndNevis.RegisterCapital(xlanguage.Japanese, "バセテール")
}
