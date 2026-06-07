package fake

import (
	xlanguage "golang.org/x/text/language"
)

// langZh is the language tag used to key Simplified Chinese name / city /
// street pools registered against [localeCN].
var langZh = xlanguage.Chinese

func init() {
	localeCN.LastNames[langZh] = chineseLastNames
	localeCN.FirstNames[langZh] = map[Gender][]string{
		GenderMale:   chineseMaleFirstNames,
		GenderFemale: chineseFemaleFirstNames,
	}
	localeCN.Cities[langZh] = chineseCities
	localeCN.Streets[langZh] = chineseStreets
}

// chineseLastNames is the pool of common Han single-character family names
// drawn from the historical 百家姓 ordering. Covers the ~120 most frequent
// surnames in mainland China per recent demographic statistics.
var chineseLastNames = []string{
	"王", "李", "张", "刘", "陈", "杨", "赵", "黄", "周", "吴",
	"徐", "孙", "胡", "朱", "高", "林", "何", "郭", "马", "罗",
	"梁", "宋", "郑", "谢", "韩", "唐", "冯", "于", "董", "萧",
	"程", "曹", "袁", "邓", "许", "傅", "沈", "曾", "彭", "吕",
	"苏", "卢", "蒋", "蔡", "贾", "丁", "魏", "薛", "叶", "阎",
	"余", "潘", "杜", "戴", "夏", "钟", "汪", "田", "任", "姜",
	"范", "方", "石", "姚", "谭", "廖", "邹", "熊", "金", "陆",
	"郝", "孔", "白", "崔", "康", "毛", "邱", "秦", "江", "史",
	"顾", "侯", "邵", "孟", "龙", "万", "段", "章", "钱", "汤",
	"尹", "黎", "易", "常", "武", "乔", "贺", "赖", "龚", "文",
	"庞", "樊", "兰", "殷", "施", "陶", "洪", "翟", "安", "颜",
	"倪", "严", "牛", "温", "芦", "季", "俞", "章", "鲁", "葛",
}

// chineseMaleFirstNames lists masculine given names, mixing common single
// characters and modern double-character pairings sampled from contemporary
// Chinese baby-naming statistics.
var chineseMaleFirstNames = []string{
	"伟", "强", "磊", "勇", "军", "峰", "刚", "平", "辉", "明",
	"建国", "建华", "建军", "建平", "建明", "建强", "建勇", "建斌", "建辉", "建生",
	"志强", "志伟", "志明", "志勇", "志刚", "志华", "志远", "志鹏", "志宏", "志超",
	"小军", "小明", "小强", "小伟", "小峰", "小刚", "小波", "小龙", "小磊", "小华",
	"俊杰", "俊伟", "俊辉", "俊豪", "俊宇", "俊熙", "俊驰", "俊楠", "俊彦", "俊凯",
	"泽宇", "泽华", "泽轩", "泽明", "泽凯", "泽涛", "泽鹏", "泽天", "泽霖", "泽阳",
	"子轩", "子涵", "子豪", "子辰", "子健", "子昂", "子谦", "子墨", "子文", "子明",
	"昊宇", "昊然", "昊轩", "昊东", "昊磊", "昊昱", "昊天", "昊辰", "昊鹏", "昊洋",
	"浩然", "浩宇", "浩轩", "浩天", "浩瀚", "浩东", "浩鹏", "浩明", "浩翔", "浩华",
	"宇轩", "宇航", "宇辰", "宇翔", "宇豪", "宇晨", "宇泽", "宇宁", "宇飞", "宇浩",
	"梓轩", "梓豪", "梓恒", "梓涵", "梓睿", "梓昂", "梓彬", "梓晨", "梓骏", "梓铭",
	"睿哲", "睿明", "睿渊", "睿轩", "睿杰", "睿涵", "睿森", "睿恒", "睿安", "睿翔",
	"哲瀚", "哲彦", "哲宇", "展鹏", "展浩", "展尚", "彦霖", "彦宇", "彦昌", "彦庭",
	"锦程", "锦堂", "锦添", "景天", "景行", "景明", "嘉懿", "嘉良", "嘉禾", "嘉慕",
	"学军", "学文", "学武", "学斌", "学伟", "学明", "国强", "国华", "国伟", "国栋",
	"明远", "明辉", "明哲", "明轩", "明宇", "明俊", "致远", "承志", "承德", "承业",
	"耀祖", "耀文", "耀辉", "耀华", "耀东", "耀彬", "海涛", "海洋", "海波", "海明",
	"云鹏", "云飞", "云山", "云逸", "云轩", "正豪", "正杰", "正阳", "正明", "正鹏",
	"凯文", "凯旋", "凯杰", "凯歌", "凯华", "凯泽", "天宇", "天翔", "天明", "天佑",
	"文博", "文昊", "文翰", "文彬", "文康", "文嘉", "文涛", "文轩", "文渊", "文杰",
	"新民", "新华", "新宇", "新峰", "新荣", "新德", "世杰", "世豪", "世明", "世昌",
	"少华", "少军", "少杰", "少康", "少棠", "少川", "少昂", "少凡", "骏", "彬",
}

// chineseFemaleFirstNames lists feminine given names mixing classical
// single-character forms with contemporary double-character pairings.
var chineseFemaleFirstNames = []string{
	"芳", "娜", "敏", "静", "丽", "颖", "燕", "霞", "玲", "婷",
	"秀英", "秀兰", "秀梅", "秀华", "秀珍", "秀芳", "秀云", "秀玲", "秀娟", "秀红",
	"春兰", "春梅", "春花", "春英", "春艳", "春霞", "春燕", "春娟", "春玲", "春红",
	"美丽", "美华", "美玉", "美琳", "美琪", "美玲", "美芳", "美英", "美兰", "美红",
	"婉莹", "婉婷", "婉清", "婉柔", "婉怡", "婉君", "婉如", "婉仪", "婉宁", "婉童",
	"欣怡", "欣然", "欣妍", "欣彤", "欣月", "欣悦", "欣慰", "欣茹", "欣琪", "欣蕊",
	"梓涵", "梓晴", "梓萱", "梓妍", "梓琳", "梓菁", "梓萌", "梓楠", "梓嫣", "梓昕",
	"雨萱", "雨晴", "雨欣", "雨彤", "雨桐", "雨涵", "雨薇", "雨柔", "雨菲", "雨欢",
	"紫萱", "紫嫣", "紫薇", "紫萌", "紫怡", "紫琼", "紫琳", "紫蝶", "紫芸", "紫菡",
	"若曦", "若萱", "若彤", "若怡", "若兰", "若菲", "若雪", "若萌", "若芸", "若颖",
	"乐怡", "乐瑶", "乐欣", "乐心", "乐天", "乐悦", "乐尧", "乐安", "乐颜", "乐韵",
	"佳怡", "佳琪", "佳音", "佳琦", "佳颖", "佳涵", "佳玥", "佳琳", "佳卉", "佳玉",
	"君雅", "君兰", "君茹", "茹云", "茹雪", "茹欣", "茹颖", "茹琳", "茹萱", "茹梦",
	"琳琅", "琳怡", "琳玉", "琳琳", "琳娜", "琳依", "琳菲", "琳蕊", "琳惜", "琳婉",
	"莉娜", "莉雪", "莉萍", "莉霞", "莉茜", "莉琴", "莉芳", "莉清", "莉颖", "莉婷",
	"婷婷", "婷玉", "婷涵", "婷茹", "婷怡", "婷兰", "婷雯", "婷雅", "婷洁", "婷云",
	"雯婷", "雯雅", "雯欣", "雯雯", "雯萱", "雯怡", "雯琪", "雯诗", "雯心", "雯静",
	"静雯", "静怡", "静思", "静琪", "静雅", "静琳", "静儿", "静兰", "静茹", "静宜",
	"茜雯", "茜文", "茜怡", "茜彤", "茜婷", "茜玉", "茜云", "茜诺", "茜玥", "茜悦",
	"莹雪", "莹颖", "莹琳", "莹琪", "莹莹", "莹晓", "莹妍", "莹欣", "莹华", "莹芳",
	"菲菲", "菲儿", "菲雪", "菲雯", "菲琳", "菲扬", "菲洋", "菲琼", "菲妍", "菲玉",
	"颖怡", "颖琪", "颖芬", "颖姗", "颖娴", "颖琳", "颖恬", "颖菲", "颖蓉", "颖薇",
	"蓉蓉", "蓉芸", "蓉云", "蓉雅", "蓉嘉", "蓉芬", "蓉莹", "蓉颖", "蓉杏", "蓉芳",
}

// chineseCities samples real Chinese cities together with their province and
// approximate centre coordinates (±0.05° tolerance). Sourced from publicly
// available geographic data (Wikipedia / National Bureau of Statistics).
var chineseCities = []CityEntry{
	{Name: "北京", Province: "北京", Lat: 39.9042, Lng: 116.4074},
	{Name: "上海", Province: "上海", Lat: 31.2304, Lng: 121.4737},
	{Name: "广州", Province: "广东", Lat: 23.1291, Lng: 113.2644},
	{Name: "深圳", Province: "广东", Lat: 22.5431, Lng: 114.0579},
	{Name: "成都", Province: "四川", Lat: 30.5728, Lng: 104.0668},
	{Name: "杭州", Province: "浙江", Lat: 30.2741, Lng: 120.1551},
	{Name: "南京", Province: "江苏", Lat: 32.0603, Lng: 118.7969},
	{Name: "武汉", Province: "湖北", Lat: 30.5928, Lng: 114.3055},
	{Name: "西安", Province: "陕西", Lat: 34.3416, Lng: 108.9398},
	{Name: "重庆", Province: "重庆", Lat: 29.5630, Lng: 106.5516},
	{Name: "天津", Province: "天津", Lat: 39.3434, Lng: 117.3616},
	{Name: "苏州", Province: "江苏", Lat: 31.2989, Lng: 120.5853},
	{Name: "青岛", Province: "山东", Lat: 36.0671, Lng: 120.3826},
	{Name: "大连", Province: "辽宁", Lat: 38.9140, Lng: 121.6147},
	{Name: "沈阳", Province: "辽宁", Lat: 41.8057, Lng: 123.4315},
	{Name: "长春", Province: "吉林", Lat: 43.8171, Lng: 125.3235},
	{Name: "哈尔滨", Province: "黑龙江", Lat: 45.8038, Lng: 126.5350},
	{Name: "济南", Province: "山东", Lat: 36.6512, Lng: 117.1201},
	{Name: "郑州", Province: "河南", Lat: 34.7466, Lng: 113.6253},
	{Name: "合肥", Province: "安徽", Lat: 31.8206, Lng: 117.2272},
	{Name: "福州", Province: "福建", Lat: 26.0745, Lng: 119.2965},
	{Name: "厦门", Province: "福建", Lat: 24.4798, Lng: 118.0894},
	{Name: "南昌", Province: "江西", Lat: 28.6820, Lng: 115.8579},
	{Name: "长沙", Province: "湖南", Lat: 28.2282, Lng: 112.9388},
	{Name: "南宁", Province: "广西", Lat: 22.8170, Lng: 108.3669},
	{Name: "海口", Province: "海南", Lat: 20.0440, Lng: 110.1999},
	{Name: "三亚", Province: "海南", Lat: 18.2528, Lng: 109.5119},
	{Name: "昆明", Province: "云南", Lat: 24.8801, Lng: 102.8329},
	{Name: "贵阳", Province: "贵州", Lat: 26.6470, Lng: 106.6302},
	{Name: "兰州", Province: "甘肃", Lat: 36.0611, Lng: 103.8343},
	{Name: "西宁", Province: "青海", Lat: 36.6172, Lng: 101.7782},
	{Name: "银川", Province: "宁夏", Lat: 38.4872, Lng: 106.2309},
	{Name: "乌鲁木齐", Province: "新疆", Lat: 43.8256, Lng: 87.6168},
	{Name: "拉萨", Province: "西藏", Lat: 29.6520, Lng: 91.1721},
	{Name: "呼和浩特", Province: "内蒙古", Lat: 40.8426, Lng: 111.7493},
	{Name: "石家庄", Province: "河北", Lat: 38.0428, Lng: 114.5149},
	{Name: "太原", Province: "山西", Lat: 37.8706, Lng: 112.5489},
	{Name: "无锡", Province: "江苏", Lat: 31.4912, Lng: 120.3119},
	{Name: "宁波", Province: "浙江", Lat: 29.8683, Lng: 121.5440},
	{Name: "温州", Province: "浙江", Lat: 27.9938, Lng: 120.6993},
	{Name: "泉州", Province: "福建", Lat: 24.8741, Lng: 118.6757},
	{Name: "东莞", Province: "广东", Lat: 23.0207, Lng: 113.7518},
	{Name: "佛山", Province: "广东", Lat: 23.0218, Lng: 113.1219},
	{Name: "中山", Province: "广东", Lat: 22.5170, Lng: 113.3927},
	{Name: "珠海", Province: "广东", Lat: 22.2710, Lng: 113.5767},
	{Name: "汕头", Province: "广东", Lat: 23.3540, Lng: 116.6820},
	{Name: "惠州", Province: "广东", Lat: 23.1118, Lng: 114.4163},
	{Name: "烟台", Province: "山东", Lat: 37.4638, Lng: 121.4479},
	{Name: "潍坊", Province: "山东", Lat: 36.7069, Lng: 119.1619},
	{Name: "淄博", Province: "山东", Lat: 36.8131, Lng: 118.0548},
	{Name: "徐州", Province: "江苏", Lat: 34.2058, Lng: 117.2837},
	{Name: "常州", Province: "江苏", Lat: 31.7727, Lng: 119.9469},
	{Name: "扬州", Province: "江苏", Lat: 32.3942, Lng: 119.4129},
	{Name: "南通", Province: "江苏", Lat: 31.9802, Lng: 120.8943},
	{Name: "盐城", Province: "江苏", Lat: 33.3477, Lng: 120.1633},
	{Name: "连云港", Province: "江苏", Lat: 34.5969, Lng: 119.2216},
}

// chineseStreets is a pool of generic street / road name templates commonly
// observed in Chinese urban address grids. Combined with city + district
// strings to form plausible full addresses.
var chineseStreets = []string{
	"长安街", "人民路", "解放路", "建设路", "和平路", "中山路", "新华路", "光明路", "文化路", "学院路",
	"迎宾大道", "世纪大道", "南京路", "北京路", "上海路", "广州路", "深圳大道", "花园路", "东风路", "西风路",
	"胜利路", "友谊路", "团结路", "奋斗路", "跃进路", "劳动路", "工人路", "农民街", "金融街", "商业街",
	"步行街", "明珠路", "龙凤街", "虎山路", "凤凰路", "梅花路", "兰花路", "樱花路", "桂花街", "玫瑰路",
	"牡丹路", "松林路", "柏树街", "槐花街", "杏林路", "桃源路", "柳荫路", "榕树街", "枫叶街", "菊花路",
	"环城路", "滨江路", "沿河街", "山前路", "湖滨路", "湖光路",
}
