package fake

import (
	xlanguage "golang.org/x/text/language"
)

// langJa is the BCP-47 tag for Japanese, used as the key when registering the
// JP locale's native data pools. Hoisted to a package-level var so callers and
// tests can reference the exact tag used during init.
var langJa = xlanguage.Japanese

// init fills the JP locale's per-language data pools with Japanese data. The
// file carries no build tag because Japanese is an official language of Japan
// and the per-entity locale convention exempts official languages so that
// localeJP always has at least one populated language pool.
func init() {
	localeJP.LastNames[langJa] = japaneseLastNames
	localeJP.FirstNames[langJa] = map[Gender][]string{
		GenderMale:   japaneseMaleFirstNames,
		GenderFemale: japaneseFemaleFirstNames,
	}
	localeJP.Cities[langJa] = japaneseCities
	localeJP.Streets[langJa] = japaneseStreets
}

// japaneseLastNames samples common Japanese family names (myoji) drawn from
// the Meiji Yasuda Life surname ranking. All entries are kanji.
var japaneseLastNames = []string{
	"佐藤", "鈴木", "高橋", "田中", "伊藤", "渡辺", "山本", "中村", "小林", "加藤",
	"吉田", "山田", "佐々木", "山口", "松本", "井上", "木村", "林", "斎藤", "清水",
	"山崎", "森", "池田", "橋本", "阿部", "石川", "前田", "藤田", "後藤", "小川",
	"岡田", "村上", "長谷川", "近藤", "石井", "坂本", "遠藤", "青木", "藤井", "西村",
	"福田", "太田", "三浦", "岡本", "松田", "中島", "原田", "小野", "田村", "竹内",
	"金子", "和田", "中山", "石田", "上田", "森田", "原", "柴田", "酒井", "工藤",
	"横山", "宮崎", "宮本", "内田", "高木", "安藤", "島田", "谷口", "大野", "高田",
	"丸山", "今井", "河野", "藤原", "村田", "武田", "上野", "杉山", "増田", "小島",
	"小山", "大塚", "平野", "菅原", "久保", "松井", "千葉", "岩崎", "桜井", "野口",
	"松尾", "野村", "渡部", "菊地", "木下", "佐野", "市川", "水野", "新井", "杉本",
}

// japaneseMaleFirstNames samples common Japanese male given names drawn from
// the Ministry of Health, Labour and Welfare's annual newborn name rankings
// and historical baby-name compilations. Entries mix kanji, hiragana, and a
// handful of katakana renderings encountered in modern usage.
var japaneseMaleFirstNames = []string{
	"大翔", "蓮", "樹", "湊", "陽翔", "悠真", "蒼", "結翔", "朝陽", "新",
	"颯", "蒼空", "凪", "奏", "結斗", "悠人", "陽斗", "湊斗", "悠斗", "碧",
	"陽向", "蒼介", "陽", "葵", "律", "颯真", "翔", "晴", "瑛太", "颯太",
	"陸", "悠", "翔太", "海斗", "悠生", "陽大", "玲央", "悠翔", "蒼太", "怜",
	"匠", "拓海", "和真", "海", "陽介", "健太", "誠", "翼", "亮", "直樹",
	"裕太", "雄大", "大樹", "翔平", "拓也", "達也", "智也", "雅人", "和也", "孝太郎",
	"亮太", "翔吾", "祐介", "啓太", "賢太", "竜也", "勇気", "光", "晃", "聡",
	"修", "学", "亮介", "雄太", "颯介", "蓮斗", "陽生", "悠太", "海翔", "新太",
	"瑛斗", "陽翔太", "結人", "悠陽", "蒼史", "蒼平", "湊太", "煌", "凛太郎", "瑞希",
	"大和", "弘樹", "智", "貴大", "雄一", "克彦", "敏夫", "正樹", "幸雄", "義明",
	"康介", "信一", "博之", "茂", "勝", "豊", "弘", "司",
}

// japaneseFemaleFirstNames samples common Japanese female given names drawn
// from the same rankings as the male set, again mixing kanji, hiragana and
// a small katakana cohort that appears in modern usage.
var japaneseFemaleFirstNames = []string{
	"陽葵", "凛", "結菜", "紬", "澪", "葵", "結愛", "莉子", "咲良", "美月",
	"ひなた", "花", "結月", "琴音", "愛子", "茉莉", "桜", "美咲", "由依", "遥",
	"心春", "結衣", "陽菜", "心愛", "詩", "杏", "美桜", "千尋", "心結", "莉央",
	"優奈", "彩花", "莉緒", "陽愛", "美羽", "百花", "美緒", "葉月", "凜花", "ひより",
	"芽依", "椿", "ことね", "美結", "あかり", "瑠奈", "和花", "栞", "美波", "理子",
	"沙織", "麻衣", "由美子", "智子", "恵子", "美穂", "綾", "美保", "由香", "里奈",
	"愛美", "彩", "千夏", "舞", "葵衣", "明日香", "美希", "真央", "由佳", "美佳",
	"千春", "佳奈", "美奈", "麻里", "亜美", "瞳", "桃子", "梨花", "茜", "彩乃",
	"優花", "結菜美", "実咲", "未来", "希", "美月香", "桜子", "杏奈", "莉奈", "愛",
	"優", "愛梨", "ゆりか", "ひかり", "りん", "ひな", "ちひろ", "あおい", "さくら", "つむぎ",
	"はるか", "みお", "みゆ", "ゆずき", "あいり", "りお", "なな", "まひろ",
}

// japaneseCities lists prefectural capitals and other major Japanese cities
// with their owning todofuken (prefecture). Coordinates are approximate
// city-centre values cross-checked against Wikipedia.
var japaneseCities = []CityEntry{
	{Name: "東京", Province: "東京都", Lat: 35.6762, Lng: 139.6503},
	{Name: "大阪", Province: "大阪府", Lat: 34.6937, Lng: 135.5023},
	{Name: "横浜", Province: "神奈川県", Lat: 35.4437, Lng: 139.6380},
	{Name: "名古屋", Province: "愛知県", Lat: 35.1815, Lng: 136.9066},
	{Name: "札幌", Province: "北海道", Lat: 43.0621, Lng: 141.3544},
	{Name: "京都", Province: "京都府", Lat: 35.0116, Lng: 135.7681},
	{Name: "神戸", Province: "兵庫県", Lat: 34.6901, Lng: 135.1955},
	{Name: "福岡", Province: "福岡県", Lat: 33.5904, Lng: 130.4017},
	{Name: "広島", Province: "広島県", Lat: 34.3853, Lng: 132.4553},
	{Name: "仙台", Province: "宮城県", Lat: 38.2682, Lng: 140.8694},
	{Name: "千葉", Province: "千葉県", Lat: 35.6074, Lng: 140.1065},
	{Name: "さいたま", Province: "埼玉県", Lat: 35.8617, Lng: 139.6455},
	{Name: "北九州", Province: "福岡県", Lat: 33.8835, Lng: 130.8752},
	{Name: "堺", Province: "大阪府", Lat: 34.5733, Lng: 135.4830},
	{Name: "新潟", Province: "新潟県", Lat: 37.9026, Lng: 139.0233},
	{Name: "浜松", Province: "静岡県", Lat: 34.7108, Lng: 137.7261},
	{Name: "熊本", Province: "熊本県", Lat: 32.8032, Lng: 130.7079},
	{Name: "相模原", Province: "神奈川県", Lat: 35.5713, Lng: 139.3729},
	{Name: "岡山", Province: "岡山県", Lat: 34.6551, Lng: 133.9195},
	{Name: "静岡", Province: "静岡県", Lat: 34.9756, Lng: 138.3828},
	{Name: "船橋", Province: "千葉県", Lat: 35.6946, Lng: 139.9826},
	{Name: "川口", Province: "埼玉県", Lat: 35.8079, Lng: 139.7242},
	{Name: "八王子", Province: "東京都", Lat: 35.6663, Lng: 139.3160},
	{Name: "姫路", Province: "兵庫県", Lat: 34.8151, Lng: 134.6857},
	{Name: "松山", Province: "愛媛県", Lat: 33.8392, Lng: 132.7657},
	{Name: "東大阪", Province: "大阪府", Lat: 34.6795, Lng: 135.6010},
	{Name: "西宮", Province: "兵庫県", Lat: 34.7373, Lng: 135.3414},
	{Name: "倉敷", Province: "岡山県", Lat: 34.5851, Lng: 133.7720},
	{Name: "福山", Province: "広島県", Lat: 34.4858, Lng: 133.3625},
	{Name: "尼崎", Province: "兵庫県", Lat: 34.7330, Lng: 135.4070},
	{Name: "長崎", Province: "長崎県", Lat: 32.7503, Lng: 129.8779},
	{Name: "町田", Province: "東京都", Lat: 35.5462, Lng: 139.4470},
	{Name: "金沢", Province: "石川県", Lat: 36.5613, Lng: 136.6562},
	{Name: "大津", Province: "滋賀県", Lat: 35.0045, Lng: 135.8686},
	{Name: "横須賀", Province: "神奈川県", Lat: 35.2815, Lng: 139.6720},
	{Name: "富山", Province: "富山県", Lat: 36.6953, Lng: 137.2114},
	{Name: "高松", Province: "香川県", Lat: 34.3401, Lng: 134.0434},
	{Name: "旭川", Province: "北海道", Lat: 43.7708, Lng: 142.3650},
	{Name: "長野", Province: "長野県", Lat: 36.6485, Lng: 138.1812},
	{Name: "那覇", Province: "沖縄県", Lat: 26.2125, Lng: 127.6792},
	{Name: "久留米", Province: "福岡県", Lat: 33.3192, Lng: 130.5083},
	{Name: "明石", Province: "兵庫県", Lat: 34.6431, Lng: 134.9974},
	{Name: "松戸", Province: "千葉県", Lat: 35.7777, Lng: 139.9032},
	{Name: "盛岡", Province: "岩手県", Lat: 39.7036, Lng: 141.1527},
	{Name: "宇都宮", Province: "栃木県", Lat: 36.5551, Lng: 139.8828},
	{Name: "前橋", Province: "群馬県", Lat: 36.3895, Lng: 139.0634},
	{Name: "水戸", Province: "茨城県", Lat: 36.3414, Lng: 140.4467},
	{Name: "甲府", Province: "山梨県", Lat: 35.6620, Lng: 138.5683},
	{Name: "鹿児島", Province: "鹿児島県", Lat: 31.5602, Lng: 130.5581},
	{Name: "宮崎", Province: "宮崎県", Lat: 31.9077, Lng: 131.4202},
}

// japaneseStreets samples common Tokyo / Osaka thoroughfares plus the
// generic "<area>通り" pattern that maps cleanly onto Japanese address
// templates.
var japaneseStreets = []string{
	"桜通り", "本町", "中央通り", "駅前通り", "銀座通り",
	"新橋通り", "表参道", "明治通り", "青山通り", "六本木通り",
	"渋谷通り", "原宿通り", "秋葉原通り", "上野通り", "浅草通り",
	"新宿通り", "池袋通り", "大手町", "丸の内", "八重洲",
	"日本橋", "築地", "品川駅前", "横浜駅西口", "元町",
	"中華街通り", "桜木町", "関内", "みなとみらい", "山下公園通り",
	"梅田", "難波", "心斎橋", "天王寺", "京橋",
	"四条通り", "河原町", "祇園", "三宮", "元町通り",
}
