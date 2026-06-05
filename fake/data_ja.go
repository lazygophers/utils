package fake

func init() {
	registerDataSet("ja", "addresses", "cities", dsJa_0)
	registerDataSet("ja", "addresses", "streets", dsJa_1)
	registerDataSet("ja", "companies", "names", dsJa_2)
	registerDataSet("ja", "companies", "suffixes", dsJa_3)
	registerDataSet("ja", "names", "first_female", dsJa_4)
	registerDataSet("ja", "names", "first_male", dsJa_5)
	registerDataSet("ja", "names", "last", dsJa_6)
	registerDataSet("ja", "texts", "lorem", dsJa_7)
}

var dsJa_0 = &DataSet{
	Language: "ja",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "東京", Weight: 2.0, Tags: []string{"major"}},
		{Value: "大阪", Weight: 1.8, Tags: []string{"major"}},
		{Value: "横浜", Weight: 1.6, Tags: []string{"major"}},
		{Value: "名古屋", Weight: 1.4, Tags: []string{"medium"}},
		{Value: "札幌", Weight: 1.2, Tags: []string{"medium"}},
		{Value: "神戸", Tags: []string{"medium"}},
		{Value: "京都", Weight: 0.7999999999999998, Tags: []string{"medium"}},
		{Value: "福岡", Weight: 0.5999999999999999, Tags: []string{"medium"}},
		{Value: "仙台", Weight: 0.5, Tags: []string{"medium"}},
		{Value: "千葉", Weight: 0.5, Tags: []string{"medium"}},
	},
}

var dsJa_1 = &DataSet{
	Language: "ja",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "中央通り", Weight: 2.0, Tags: []string{"common"}},
		{Value: "駅前通り", Weight: 1.7, Tags: []string{"common"}},
		{Value: "本町", Weight: 1.4, Tags: []string{"common"}},
		{Value: "大通り", Weight: 1.1, Tags: []string{"common"}},
		{Value: "商店街", Weight: 0.8, Tags: []string{"common"}},
	},
}

var dsJa_2 = &DataSet{
	Language: "ja",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "トヨタ", Weight: 2.0, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "ソニー", Weight: 1.85, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "任天堂", Weight: 1.7, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "ソフトバンク", Weight: 1.55, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "三菱", Weight: 1.4, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "ホンダ", Weight: 1.25, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "キヤノン", Weight: 1.1, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "パナソニック", Weight: 0.95, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "日産", Weight: 0.8, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "富士通", Weight: 0.6500000000000001, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
	},
}

var dsJa_3 = &DataSet{
	Language: "ja",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "株式会社", Weight: 3.0, Tags: []string{"common"}},
		{Value: "有限会社", Weight: 2.7, Tags: []string{"common"}},
		{Value: "合同会社", Weight: 2.4, Tags: []string{"common"}},
		{Value: "合資会社", Weight: 2.1, Tags: []string{"formal"}},
		{Value: "企業", Weight: 1.8, Tags: []string{"formal"}},
		{Value: "グループ", Weight: 1.5, Tags: []string{"formal"}},
	},
}

var dsJa_4 = &DataSet{
	Language: "ja",
	Country:  "JP",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "花子", Weight: 2.0, Tags: []string{"traditional", "flower"}},
		{Value: "さくら", Weight: 1.8, Tags: []string{"traditional", "cherry"}},
		{Value: "美香", Weight: 1.6, Tags: []string{"traditional", "beauty"}},
		{Value: "ゆい", Weight: 1.4, Tags: []string{"modern", "tie"}},
		{Value: "あかり", Weight: 1.2, Tags: []string{"modern", "light"}},
		{Value: "みお", Weight: 1.1, Tags: []string{"modern", "beautiful"}},
		{Value: "りお", Tags: []string{"modern", "jasmine"}},
		{Value: "あおい", Weight: 0.9, Tags: []string{"modern", "blue"}},
		{Value: "えま", Weight: 0.8, Tags: []string{"modern", "love"}},
		{Value: "はな", Weight: 0.7, Tags: []string{"traditional", "flower"}},
	},
}

var dsJa_5 = &DataSet{
	Language: "ja",
	Country:  "JP",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "太郎", Weight: 2.0, Tags: []string{"traditional", "eldest"}},
		{Value: "一郎", Weight: 1.8, Tags: []string{"traditional", "first"}},
		{Value: "二郎", Weight: 1.6, Tags: []string{"traditional", "second"}},
		{Value: "健太", Weight: 1.4, Tags: []string{"modern", "health"}},
		{Value: "翔太", Weight: 1.2, Tags: []string{"modern", "soar"}},
		{Value: "ひろし", Weight: 1.1, Tags: []string{"traditional", "wide"}},
		{Value: "まさき", Tags: []string{"traditional", "correct"}},
		{Value: "たくや", Weight: 0.9, Tags: []string{"modern", "expand"}},
		{Value: "ゆうき", Weight: 0.8, Tags: []string{"modern", "courage"}},
		{Value: "りょう", Weight: 0.7, Tags: []string{"modern", "cool"}},
	},
}

var dsJa_6 = &DataSet{
	Language: "ja",
	Country:  "JP",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "佐藤", Weight: 2.0, Tags: []string{"geographical"}},
		{Value: "鈴木", Weight: 1.8, Tags: []string{"nature"}},
		{Value: "高橋", Weight: 1.6, Tags: []string{"geographical"}},
		{Value: "田中", Weight: 1.4, Tags: []string{"geographical"}},
		{Value: "渡辺", Weight: 1.2, Tags: []string{"geographical"}},
		{Value: "伊藤", Weight: 1.1, Tags: []string{"geographical"}},
		{Value: "山本", Tags: []string{"geographical"}},
		{Value: "中村", Weight: 0.9, Tags: []string{"geographical"}},
		{Value: "小林", Weight: 0.8, Tags: []string{"geographical"}},
		{Value: "加藤", Weight: 0.7, Tags: []string{"geographical"}},
	},
}

var dsJa_7 = &DataSet{
	Language: "ja",
	Type:     "texts",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "テキスト"},
		{Value: "文字", Weight: 0.95},
		{Value: "文章", Weight: 0.9},
		{Value: "段落", Weight: 0.85},
		{Value: "記事", Weight: 0.8},
		{Value: "報告", Weight: 0.75},
		{Value: "研究", Weight: 0.7},
		{Value: "分析", Weight: 0.6499999999999999},
		{Value: "内容", Weight: 0.6},
		{Value: "話題", Weight: 0.55},
		{Value: "概念", Weight: 0.5},
		{Value: "意味", Weight: 0.44999999999999996},
		{Value: "言語", Weight: 0.3999999999999999},
		{Value: "文学", Weight: 0.35},
		{Value: "書き物", Weight: 0.29999999999999993},
	},
}
