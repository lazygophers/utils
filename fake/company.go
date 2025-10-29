package fake

import (
	"fmt"
	"strings"

	"github.com/lazygophers/utils/randx"
)

// Company 公司信息结构体
type Company struct {
	Name        string `json:"name"`
	Industry    string `json:"industry"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Founded     int    `json:"founded"`
	Employees   int    `json:"employees"`
}

// CompanyName 生成公司名称
func (f *Faker) CompanyName() string {

	// 尝试从数据文件加载
	names, weights, err := getWeightedItems(f.language, "companies", "names")
	if err != nil {
		// 回退到生成模式
		return f.generateCompanyName()
	}

	return randx.WeightedChoose(names, weights)
}

// generateCompanyName 生成式公司名称
func (f *Faker) generateCompanyName() string {
	patterns := []string{
		"%s %s",    // Apple Inc.
		"%s & %s",  // Smith & Johnson
		"%s-%s",    // Micro-Systems
		"%s%s",     // TechCorp
		"%s %s %s", // Global Tech Solutions
	}

	firstNames := f.getCompanyFirstNames()
	suffixes := f.getCompanySuffixes()

	pattern := randx.Choose(patterns)

	switch pattern {
	case "%s %s":
		firstName := randx.Choose(firstNames)
		suffix := randx.Choose(suffixes)
		return fmt.Sprintf(pattern, firstName, suffix)

	case "%s & %s":
		name1 := randx.Choose(firstNames)
		name2 := randx.Choose(firstNames)
		suffix := randx.Choose(suffixes)
		return fmt.Sprintf("%s & %s %s", name1, name2, suffix)

	case "%s-%s":
		name1 := randx.Choose(firstNames)
		name2 := randx.Choose(firstNames)
		return fmt.Sprintf(pattern, name1, name2)

	case "%s%s":
		name1 := randx.Choose(firstNames)
		name2 := randx.Choose(firstNames)
		return fmt.Sprintf(pattern, name1, name2)

	case "%s %s %s":
		adj := randx.Choose([]string{"Global", "Advanced", "Premium", "Elite", "Smart", "Digital", "Modern", "Future", "Dynamic", "Innovative"})
		name := randx.Choose(firstNames)
		suffix := randx.Choose(suffixes)
		return fmt.Sprintf(pattern, adj, name, suffix)

	default:
		firstName := randx.Choose(firstNames)
		suffix := randx.Choose(suffixes)
		return fmt.Sprintf("%s %s", firstName, suffix)
	}
}

func (f *Faker) getCompanyFirstNames() []string {
	return []string{
		"Tech", "Data", "Cloud", "Digital", "Smart", "Quantum", "Cyber", "Global",
		"Alpha", "Beta", "Gamma", "Delta", "Sigma", "Omega", "Nova", "Apex",
		"Prime", "Elite", "Pro", "Max", "Ultra", "Super", "Mega", "Hyper",
		"Micro", "Nano", "Meta", "Neo", "Auto", "Eco", "Bio", "Geo",
		"Aero", "Hydro", "Electro", "Thermo", "Solar", "Lunar", "Stellar",
		"Universal", "National", "International", "Regional", "Local", "United",
		"Advanced", "Modern", "Future", "Next", "New", "Dynamic", "Rapid",
		"Swift", "Flash", "Bolt", "Lightning", "Thunder", "Storm", "Wind",
		"Ocean", "River", "Mountain", "Valley", "Forest", "Desert", "Arctic",
		"Phoenix", "Eagle", "Lion", "Tiger", "Wolf", "Bear", "Shark", "Falcon",
		"Diamond", "Gold", "Silver", "Platinum", "Crystal", "Pearl", "Ruby",
		"Blue", "Red", "Green", "Yellow", "Purple", "Orange", "Black", "White",
	}
}

func (f *Faker) getCompanySuffixes() []string {
	// 尝试从数据文件加载后缀
	suffixes, _, err := getWeightedItems(f.language, "companies", "suffixes")
	if err == nil {
		return suffixes
	}

	// 回退默认值
	return []string{
		"Inc.", "LLC", "Corp.", "Co.", "Ltd.", "Group", "Holdings",
		"Company", "Corporation", "Enterprises", "Industries", "Systems",
		"Solutions", "Technologies", "Services", "Associates", "Partners",
		"Consulting", "Global", "International", "Worldwide", "National",
	}
}

// CompanySuffix 生成公司后缀
func (f *Faker) CompanySuffix() string {

	suffixes, weights, err := getWeightedItems(f.language, "companies", "suffixes")
	if err != nil {
		suffixes = f.getCompanySuffixes()
		return randx.Choose(suffixes)
	}

	return randx.WeightedChoose(suffixes, weights)
}

// Industry 生成行业名称
func (f *Faker) Industry() string {

	industries := []string{
		"Technology", "Healthcare", "Finance", "Education", "Manufacturing",
		"Retail", "Hospitality", "Transportation", "Energy", "Construction",
		"Agriculture", "Entertainment", "Media", "Telecommunications", "Insurance",
		"Real Estate", "Automotive", "Aerospace", "Biotechnology", "Pharmaceuticals",
		"Software", "Hardware", "Consulting", "Marketing", "Advertising",
		"E-commerce", "Logistics", "Supply Chain", "Banking", "Investment",
		"Food & Beverage", "Fashion", "Beauty", "Sports", "Gaming",
		"Publishing", "Legal", "Architecture", "Engineering", "Research",
		"Non-profit", "Government", "Military", "Space", "Ocean",
		"Environmental", "Renewable Energy", "Mining", "Oil & Gas", "Utilities",
	}

	return randx.Choose(industries)
}

// JobTitle 生成职位名称
func (f *Faker) JobTitle() string {

	// 职位级别
	levels := []string{"", "Junior", "Senior", "Lead", "Principal", "Staff", "Chief"}
	level := randx.Choose(levels)

	// 职位类型
	positions := []string{
		"Engineer", "Developer", "Designer", "Manager", "Director", "Analyst",
		"Consultant", "Specialist", "Coordinator", "Administrator", "Executive",
		"Officer", "Representative", "Associate", "Assistant", "Supervisor",
		"Technician", "Architect", "Scientist", "Researcher", "Strategist",
		"Planner", "Advisor", "Expert", "Leader", "Head", "Vice President",
		"President", "CEO", "CTO", "CIO", "CFO", "CMO", "CHRO", "COO",
	}
	position := randx.Choose(positions)

	// 专业领域
	fields := []string{
		"Software", "Hardware", "Network", "Security", "Data", "AI", "ML",
		"Marketing", "Sales", "Finance", "HR", "Operations", "Product",
		"Business", "Strategy", "Design", "UX", "UI", "Frontend", "Backend",
		"DevOps", "QA", "Testing", "Support", "Customer", "Technical",
		"Project", "Program", "Portfolio", "Risk", "Compliance", "Legal",
	}

	// 30% 概率添加专业领域
	if randx.Float32() < 0.3 {
		field := randx.Choose(fields)
		if level != "" {
			return fmt.Sprintf("%s %s %s", level, field, position)
		}
		return fmt.Sprintf("%s %s", field, position)
	}

	if level != "" {
		return fmt.Sprintf("%s %s", level, position)
	}

	return position
}

// Department 生成部门名称
func (f *Faker) Department() string {

	departments := []string{
		"Engineering", "Development", "Design", "Product", "Marketing", "Sales",
		"Finance", "Accounting", "Human Resources", "Operations", "Support",
		"Customer Service", "Business Development", "Strategy", "Research",
		"Quality Assurance", "Testing", "DevOps", "Infrastructure", "Security",
		"Data", "Analytics", "AI", "Machine Learning", "Legal", "Compliance",
		"Executive", "Administration", "Facilities", "IT", "Procurement",
		"Supply Chain", "Logistics", "Manufacturing", "Production", "Warehouse",
		"Shipping", "Receiving", "Inventory", "Planning", "Forecasting",
		"Risk Management", "Audit", "Training", "Learning", "Innovation",
	}

	return randx.Choose(departments)
}

// CompanyInfo 生成完整公司信息
func (f *Faker) CompanyInfo() *Company {

	name := f.CompanyName()
	industry := f.Industry()

	// 生成公司描述
	descriptions := []string{
		"Leading provider of innovative %s solutions",
		"Global leader in %s technology",
		"Premier %s company serving worldwide clients",
		"Cutting-edge %s solutions for modern businesses",
		"Trusted %s partner for enterprise clients",
		"Revolutionary %s platform changing the industry",
		"Expert %s consultancy with proven track record",
		"Next-generation %s company driving innovation",
		"Award-winning %s solutions provider",
		"Comprehensive %s services for all business needs",
	}
	descTemplate := randx.Choose(descriptions)
	description := fmt.Sprintf(descTemplate, strings.ToLower(industry))

	// 生成网站域名
	domain := strings.ToLower(strings.ReplaceAll(name, " ", ""))
	domain = strings.ReplaceAll(domain, ".", "")
	domain = strings.ReplaceAll(domain, "&", "and")
	domain = strings.ReplaceAll(domain, "-", "")
	website := fmt.Sprintf("https://www.%s.com", domain)

	// 生成公司邮箱
	emailPrefix := []string{"info", "contact", "hello", "support", "sales", "admin"}
	email := fmt.Sprintf("%s@%s.com", randx.Choose(emailPrefix), domain)

	// 生成成立年份
	currentYear := 2024
	founded := randx.Intn(currentYear-1900) + 1900

	// 生成员工数量
	employeeRanges := [][]int{
		{1, 10},         // Startup
		{11, 50},        // Small
		{51, 200},       // Medium
		{201, 1000},     // Large
		{1001, 10000},   // Enterprise
		{10001, 100000}, // Multinational
	}
	employeeRange := randx.Choose(employeeRanges)
	employees := randx.Intn(employeeRange[1]-employeeRange[0]+1) + employeeRange[0]

	return &Company{
		Name:        name,
		Industry:    industry,
		Description: description,
		Website:     website,
		Email:       email,
		Phone:       f.PhoneNumber(),
		Address:     f.AddressLine(),
		Founded:     founded,
		Employees:   employees,
	}
}

// BS 生成商业术语/废话（Business Speak）
func (f *Faker) BS() string {

	verbs := []string{
		"implement", "utilize", "integrate", "streamline", "optimize", "leverage",
		"enhance", "facilitate", "maximize", "revolutionize", "transform", "innovate",
		"deploy", "architect", "engineer", "develop", "create", "build",
		"deliver", "provide", "enable", "empower", "accelerate", "scale",
		"modernize", "digitize", "automate", "orchestrate", "synchronize", "align",
	}

	adjectives := []string{
		"strategic", "tactical", "dynamic", "robust", "scalable", "flexible",
		"innovative", "cutting-edge", "next-generation", "world-class", "best-in-class",
		"enterprise-grade", "mission-critical", "high-performance", "real-time",
		"cross-platform", "end-to-end", "seamless", "integrated", "comprehensive",
		"customizable", "user-friendly", "intuitive", "efficient", "effective",
	}

	nouns := []string{
		"solutions", "systems", "platforms", "frameworks", "architectures", "infrastructures",
		"applications", "services", "technologies", "methodologies", "processes", "workflows",
		"strategies", "initiatives", "programs", "projects", "deliverables", "outcomes",
		"synergies", "paradigms", "ecosystems", "networks", "channels", "interfaces",
		"functionalities", "capabilities", "features", "components", "modules", "elements",
	}

	verb := randx.Choose(verbs)
	adjective := randx.Choose(adjectives)
	noun := randx.Choose(nouns)

	patterns := []string{
		"%s %s %s",
		"%s world-class %s",
		"efficiently %s %s %s",
		"seamlessly %s next-generation %s",
		"globally %s enterprise-wide %s",
	}

	pattern := randx.Choose(patterns)

	switch pattern {
	case "%s %s %s":
		return fmt.Sprintf(pattern, verb, adjective, noun)
	case "%s world-class %s":
		return fmt.Sprintf(pattern, verb, noun)
	case "efficiently %s %s %s":
		return fmt.Sprintf(pattern, verb, adjective, noun)
	case "seamlessly %s next-generation %s":
		return fmt.Sprintf(pattern, verb, noun)
	case "globally %s enterprise-wide %s":
		return fmt.Sprintf(pattern, verb, noun)
	default:
		return fmt.Sprintf("%s %s %s", verb, adjective, noun)
	}
}

// Catchphrase 生成公司口号
func (f *Faker) Catchphrase() string {

	patterns := []string{
		"Think %s, Act %s",
		"%s Solutions for a %s World",
		"Where %s Meets %s",
		"Building the %s of Tomorrow",
		"Your %s Partner for %s",
		"Empowering %s Through %s",
		"Leading the %s Revolution",
		"Innovation in %s, Excellence in %s",
	}

	adjectives := []string{
		"Smart", "Digital", "Global", "Future", "Innovative", "Dynamic", "Strategic",
		"Advanced", "Modern", "Progressive", "Cutting-edge", "Revolutionary",
	}

	nouns := []string{
		"Technology", "Innovation", "Excellence", "Quality", "Performance", "Growth",
		"Success", "Solutions", "Results", "Value", "Trust", "Partnership",
	}

	pattern := randx.Choose(patterns)

	switch pattern {
	case "Think %s, Act %s":
		return fmt.Sprintf(pattern, randx.Choose(adjectives), randx.Choose(adjectives))
	case "%s Solutions for a %s World":
		return fmt.Sprintf(pattern, randx.Choose(adjectives), randx.Choose(adjectives))
	case "Where %s Meets %s":
		return fmt.Sprintf(pattern, randx.Choose(nouns), randx.Choose(nouns))
	case "Building the %s of Tomorrow":
		return fmt.Sprintf(pattern, randx.Choose(nouns))
	case "Your %s Partner for %s":
		return fmt.Sprintf(pattern, randx.Choose(adjectives), randx.Choose(nouns))
	case "Empowering %s Through %s":
		return fmt.Sprintf(pattern, randx.Choose(nouns), randx.Choose(nouns))
	case "Leading the %s Revolution":
		return fmt.Sprintf(pattern, randx.Choose(nouns))
	case "Innovation in %s, Excellence in %s":
		return fmt.Sprintf(pattern, randx.Choose(nouns), randx.Choose(nouns))
	default:
		return fmt.Sprintf("Your Trusted %s Partner", randx.Choose(adjectives))
	}
}

// 批量生成函数
func (f *Faker) BatchCompanyNames(count int) []string {
	return f.batchGenerate(count, f.CompanyName)
}

func (f *Faker) BatchJobTitles(count int) []string {
	return f.batchGenerate(count, f.JobTitle)
}

func (f *Faker) BatchCompanyInfos(count int) []*Company {
	results := make([]*Company, count)
	for i := 0; i < count; i++ {
		results[i] = f.CompanyInfo()
	}
	return results
}

// 全局便捷函数
func CompanyName() string {
	return getDefaultFaker().CompanyName()
}

func CompanySuffix() string {
	return getDefaultFaker().CompanySuffix()
}

func Industry() string {
	return getDefaultFaker().Industry()
}

func JobTitle() string {
	return getDefaultFaker().JobTitle()
}

func Department() string {
	return getDefaultFaker().Department()
}

func CompanyInfo() *Company {
	return getDefaultFaker().CompanyInfo()
}

func BS() string {
	return getDefaultFaker().BS()
}

func Catchphrase() string {
	return getDefaultFaker().Catchphrase()
}
