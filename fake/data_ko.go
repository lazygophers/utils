package fake

func init() {
	registerDataSet("ko", "addresses", "cities", dsKo_0)
	registerDataSet("ko", "addresses", "streets", dsKo_1)
	registerDataSet("ko", "companies", "names", dsKo_2)
	registerDataSet("ko", "companies", "suffixes", dsKo_3)
	registerDataSet("ko", "names", "first_female", dsKo_4)
	registerDataSet("ko", "names", "first_male", dsKo_5)
	registerDataSet("ko", "names", "last", dsKo_6)
	registerDataSet("ko", "texts", "lorem", dsKo_7)
}

var dsKo_0 = &DataSet{
	Language: "ko",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "서울", Weight: 2.0, Tags: []string{"major"}},
		{Value: "부산", Weight: 1.8, Tags: []string{"major"}},
		{Value: "대구", Weight: 1.6, Tags: []string{"major"}},
		{Value: "인천", Weight: 1.4, Tags: []string{"medium"}},
		{Value: "광주", Weight: 1.2, Tags: []string{"medium"}},
		{Value: "대전", Tags: []string{"medium"}},
		{Value: "울산", Weight: 0.7999999999999998, Tags: []string{"medium"}},
		{Value: "창원", Weight: 0.5999999999999999, Tags: []string{"medium"}},
		{Value: "고양", Weight: 0.5, Tags: []string{"medium"}},
		{Value: "용인", Weight: 0.5, Tags: []string{"medium"}},
	},
}

var dsKo_1 = &DataSet{
	Language: "ko",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "중앙대로", Weight: 2.0, Tags: []string{"common"}},
		{Value: "역앞로", Weight: 1.7, Tags: []string{"common"}},
		{Value: "시청로", Weight: 1.4, Tags: []string{"common"}},
		{Value: "대학로", Weight: 1.1, Tags: []string{"common"}},
		{Value: "상가로", Weight: 0.8, Tags: []string{"common"}},
	},
}

var dsKo_2 = &DataSet{
	Language: "ko",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "삼성", Weight: 2.0, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "현대", Weight: 1.85, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "LG", Weight: 1.7, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "SK", Weight: 1.55, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "롯데", Weight: 1.4, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "포스코", Weight: 1.25, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "한화", Weight: 1.1, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "두산", Weight: 0.95, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "GS", Weight: 0.8, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "신세계", Weight: 0.6500000000000001, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
	},
}

var dsKo_3 = &DataSet{
	Language: "ko",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "주식회사", Weight: 3.0, Tags: []string{"common"}},
		{Value: "유한회사", Weight: 2.7, Tags: []string{"common"}},
		{Value: "합자회사", Weight: 2.4, Tags: []string{"common"}},
		{Value: "기업", Weight: 2.1, Tags: []string{"formal"}},
		{Value: "그룹", Weight: 1.8, Tags: []string{"formal"}},
		{Value: "코퍼레이션", Weight: 1.5, Tags: []string{"formal"}},
	},
}

var dsKo_4 = &DataSet{
	Language: "ko",
	Country:  "KR",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "영희", Weight: 2.0, Tags: []string{"traditional", "flower"}},
		{Value: "수진", Weight: 1.8, Tags: []string{"traditional", "excellence"}},
		{Value: "은지", Weight: 1.6, Tags: []string{"modern", "silver"}},
		{Value: "지영", Weight: 1.4, Tags: []string{"modern", "wisdom"}},
		{Value: "혜진", Weight: 1.2, Tags: []string{"traditional", "wisdom"}},
		{Value: "미경", Weight: 1.1, Tags: []string{"traditional", "beauty"}},
		{Value: "소영", Tags: []string{"modern", "small"}},
		{Value: "정아", Weight: 0.9, Tags: []string{"traditional", "correct"}},
		{Value: "유진", Weight: 0.8, Tags: []string{"modern", "flow"}},
		{Value: "민지", Weight: 0.7, Tags: []string{"modern", "people"}},
	},
}

var dsKo_5 = &DataSet{
	Language: "ko",
	Country:  "KR",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "민수", Weight: 2.0, Tags: []string{"traditional", "people"}},
		{Value: "철수", Weight: 1.8, Tags: []string{"traditional", "iron"}},
		{Value: "지훈", Weight: 1.6, Tags: []string{"modern", "wisdom"}},
		{Value: "현우", Weight: 1.4, Tags: []string{"modern", "virtuous"}},
		{Value: "태현", Weight: 1.2, Tags: []string{"traditional", "great"}},
		{Value: "준호", Weight: 1.1, Tags: []string{"modern", "standard"}},
		{Value: "성민", Tags: []string{"traditional", "accomplish"}},
		{Value: "동현", Weight: 0.9, Tags: []string{"traditional", "east"}},
		{Value: "상훈", Weight: 0.8, Tags: []string{"traditional", "reward"}},
		{Value: "정우", Weight: 0.7, Tags: []string{"traditional", "correct"}},
	},
}

var dsKo_6 = &DataSet{
	Language: "ko",
	Country:  "KR",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "김", Weight: 2.0, Tags: []string{"clan"}},
		{Value: "이", Weight: 1.8, Tags: []string{"clan"}},
		{Value: "박", Weight: 1.6, Tags: []string{"clan"}},
		{Value: "최", Weight: 1.4, Tags: []string{"clan"}},
		{Value: "정", Weight: 1.2, Tags: []string{"clan"}},
		{Value: "강", Weight: 1.1, Tags: []string{"clan"}},
		{Value: "조", Tags: []string{"clan"}},
		{Value: "윤", Weight: 0.9, Tags: []string{"clan"}},
		{Value: "장", Weight: 0.8, Tags: []string{"clan"}},
		{Value: "임", Weight: 0.7, Tags: []string{"clan"}},
	},
}

var dsKo_7 = &DataSet{
	Language: "ko",
	Type:     "texts",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "텍스트"},
		{Value: "글자", Weight: 0.95},
		{Value: "문장", Weight: 0.9},
		{Value: "단락", Weight: 0.85},
		{Value: "기사", Weight: 0.8},
		{Value: "보고서", Weight: 0.75},
		{Value: "연구", Weight: 0.7},
		{Value: "분석", Weight: 0.6499999999999999},
		{Value: "내용", Weight: 0.6},
		{Value: "주제", Weight: 0.55},
		{Value: "개념", Weight: 0.5},
		{Value: "의미", Weight: 0.44999999999999996},
		{Value: "언어", Weight: 0.3999999999999999},
		{Value: "문학", Weight: 0.35},
		{Value: "글쓰기", Weight: 0.29999999999999993},
	},
}
