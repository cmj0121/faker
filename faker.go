package faker

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

const (
	FAKE_MAX_SIZE = 16
	FAKE_IGNORE   = "-"

	FAKE_KEY         = "faker"
	FAKE_VALUE_NAME  = "name"
	FAKE_VALUE_EMAIL = "email"
)

var (
	FAKE_NAME_POOL   = []byte("abcdefghijklmnopqrstuvwxyz")
	FAKE_DOMAIN_POOL = []string{
		"com",
		"uk",
		"abb",
		"abbott",
		"abogado",
		"ac",
		"academy",
		"accenture",
		"accountant",
		"accountants",
		"active",
		"actor",
		"ad",
		"ads",
		"adult",
		"ae",
		"aeg",
		"aero",
		"af",
		"afl",
		"ag",
		"agency",
		"ai",
		"aig",
		"airforce",
		"airtel",
		"al",
		"allfinanz",
		"alsace",
		"am",
		"amsterdam",
		"an",
		"android",
		"ao",
		"apartments",
		"app",
		"aq",
		"aquarelle",
		"ar",
		"archi",
		"army",
		"arpa",
		"as",
		"asia",
		"associates",
		"at",
		"attorney",
		"au",
		"auction",
		"audio",
		"auto",
		"autos",
		"aw",
		"ax",
		"axa",
		"az",
		"azure",
		"ba",
		"band",
		"bank",
		"bar",
		"barcelona",
		"barclaycard",
		"barclays",
		"bargains",
		"bauhaus",
		"bayern",
		"bb",
		"bbc",
		"bbva",
		"bcn",
		"bd",
		"be",
		"beer",
		"bentley",
		"berlin",
		"best",
		"bf",
		"bg",
		"bh",
		"bharti",
		"bi",
		"bible",
		"bid",
		"bike",
		"bing",
		"bingo",
		"bio",
		"biz",
		"bj",
		"bl",
		"black",
		"blackfriday",
		"bloomberg",
		"blue",
		"bm",
		"bmw",
		"bn",
		"bnl",
		"bnpparibas",
		"bo",
		"boats",
		"bond",
		"boo",
		"boutique",
		"bq",
		"br",
		"bradesco",
		"bridgestone",
		"broker",
		"brother",
		"brussels",
		"bs",
		"bt",
		"budapest",
		"build",
		"builders",
		"business",
		"buzz",
		"bv",
		"bw",
		"by",
		"bz",
		"bzh",
		"ca",
		"cab",
		"cafe",
		"cal",
		"camera",
		"camp",
		"cancerresearch",
		"canon",
		"capetown",
		"capital",
		"caravan",
		"cards",
		"care",
		"career",
		"careers",
		"cars",
		"cartier",
		"casa",
		"cash",
		"casino",
		"cat",
		"catering",
		"cba",
		"cbn",
		"cc",
		"cd",
		"center",
		"ceo",
		"cern",
		"cf",
		"cfa",
		"cfd",
		"cg",
		"ch",
		"channel",
		"chat",
		"cheap",
		"chloe",
		"christmas",
		"chrome",
		"church",
		"ci",
		"cisco",
		"citic",
		"city",
		"ck",
		"cl",
		"claims",
		"cleaning",
		"click",
		"clinic",
		"clothing",
		"cloud",
		"club",
		"cm",
		"cn",
		"co",
		"coach",
		"codes",
		"coffee",
		"college",
		"cologne",
		"com",
		"commbank",
		"community",
		"company",
		"computer",
		"condos",
		"construction",
		"consulting",
		"contractors",
		"cooking",
		"cool",
		"coop",
		"corsica",
		"country",
		"coupons",
		"courses",
		"cr",
		"credit",
		"creditcard",
		"cricket",
		"crown",
		"crs",
		"cruises",
		"cu",
		"cuisinella",
		"cv",
		"cw",
		"cx",
		"cy",
		"cymru",
		"cyou",
		"cz",
		"dabur",
		"dad",
		"dance",
		"date",
		"dating",
		"datsun",
		"day",
		"dclk",
		"de",
		"deals",
		"degree",
		"delivery",
		"delta",
		"democrat",
		"dental",
		"dentist",
		"desi",
		"design",
		"dev",
		"diamonds",
		"diet",
		"digital",
		"direct",
		"directory",
		"discount",
		"dj",
		"dk",
		"dm",
		"dnp",
		"do",
		"docs",
		"dog",
		"doha",
		"domains",
		"doosan",
		"download",
		"drive",
		"durban",
		"dvag",
		"dz",
		"earth",
		"eat",
		"ec",
		"edu",
		"education",
		"ee",
		"eg",
		"eh",
		"email",
		"emerck",
		"energy",
		"engineer",
		"engineering",
		"enterprises",
		"epson",
		"equipment",
		"er",
		"erni",
		"es",
		"esq",
		"estate",
		"et",
		"eu",
		"eurovision",
		"eus",
		"events",
		"everbank",
		"exchange",
		"expert",
		"exposed",
		"express",
		"fail",
		"faith",
		"fan",
		"fans",
		"farm",
		"fashion",
		"feedback",
		"fi",
		"film",
		"finance",
		"financial",
		"firmdale",
		"fish",
		"fishing",
		"fit",
		"fitness",
		"fj",
		"fk",
		"flights",
		"florist",
		"flowers",
		"flsmidth",
		"fly",
		"fm",
		"fo",
		"foo",
		"football",
		"forex",
		"forsale",
		"forum",
		"foundation",
		"fr",
		"frl",
		"frogans",
		"fund",
		"furniture",
		"futbol",
		"fyi",
		"ga",
		"gal",
		"gallery",
		"game",
		"garden",
		"gb",
		"gbiz",
		"gd",
		"gdn",
		"ge",
		"gent",
		"genting",
		"gf",
		"gg",
		"ggee",
		"gh",
		"gi",
		"gift",
		"gifts",
		"gives",
		"gl",
		"glass",
		"gle",
		"global",
		"globo",
		"gm",
		"gmail",
		"gmo",
		"gmx",
		"gn",
		"gold",
		"goldpoint",
		"golf",
		"goo",
		"goog",
		"google",
		"gop",
		"gov",
		"gp",
		"gq",
		"gr",
		"graphics",
		"gratis",
		"green",
		"gripe",
		"gs",
		"gt",
		"gu",
		"guge",
		"guide",
		"guitars",
		"guru",
		"gw",
		"gy",
		"hamburg",
		"hangout",
		"haus",
		"healthcare",
		"help",
		"here",
		"hermes",
		"hiphop",
		"hitachi",
		"hiv",
		"hk",
		"hm",
		"hn",
		"hockey",
		"holdings",
		"holiday",
		"homedepot",
		"homes",
		"honda",
		"horse",
		"host",
		"hosting",
		"hoteles",
		"hotmail",
		"house",
		"how",
		"hr",
		"hsbc",
		"ht",
		"hu",
		"ibm",
		"icbc",
		"icu",
		"id",
		"ie",
		"ifm",
		"iinet",
		"il",
		"im",
		"immo",
		"immobilien",
		"in",
		"industries",
		"infiniti",
		"info",
		"ing",
		"ink",
		"institute",
		"insure",
		"int",
		"international",
		"investments",
		"io",
		"iq",
		"ir",
		"irish",
		"is",
		"ist",
		"istanbul",
		"it",
		"iwc",
		"java",
		"jcb",
		"je",
		"jetzt",
		"jewelry",
		"jlc",
		"jll",
		"jm",
		"jo",
		"jobs",
		"joburg",
		"jp",
		"jprs",
		"juegos",
		"kaufen",
		"kddi",
		"ke",
		"kg",
		"kh",
		"ki",
		"kim",
		"kitchen",
		"kiwi",
		"km",
		"kn",
		"koeln",
		"komatsu",
		"kp",
		"kr",
		"krd",
		"kred",
		"kw",
		"ky",
		"kyoto",
		"kz",
		"la",
		"lacaixa",
		"lancaster",
		"land",
		"lasalle",
		"lat",
		"latrobe",
		"law",
		"lawyer",
		"lb",
		"lc",
		"lds",
		"lease",
		"leclerc",
		"legal",
		"lgbt",
		"li",
		"liaison",
		"lidl",
		"life",
		"lighting",
		"limited",
		"limo",
		"link",
		"live",
		"lk",
		"loan",
		"loans",
		"lol",
		"london",
		"lotte",
		"lotto",
		"love",
		"lr",
		"ls",
		"lt",
		"ltda",
		"lu",
		"lupin",
		"luxe",
		"luxury",
		"lv",
		"ly",
		"ma",
		"madrid",
		"maif",
		"maison",
		"management",
		"mango",
		"market",
		"marketing",
		"markets",
		"marriott",
		"mba",
		"mc",
		"md",
		"me",
		"media",
		"meet",
		"melbourne",
		"meme",
		"memorial",
		"men",
		"menu",
		"mf",
		"mg",
		"mh",
		"miami",
		"microsoft",
		"mil",
		"mini",
		"mk",
		"ml",
		"mm",
		"mma",
		"mn",
		"mo",
		"mobi",
		"moda",
		"moe",
		"monash",
		"money",
		"montblanc",
		"mormon",
		"mortgage",
		"moscow",
		"motorcycles",
		"mov",
		"movie",
		"movistar",
		"mp",
		"mq",
		"mr",
		"ms",
		"mt",
		"mtn",
		"mtpc",
		"mu",
		"museum",
		"mv",
		"mw",
		"mx",
		"my",
		"mz",
		"na",
		"nadex",
		"nagoya",
		"name",
		"navy",
		"nc",
		"ne",
		"nec",
		"net",
		"netbank",
		"network",
		"neustar",
		"new",
		"news",
		"nexus",
		"nf",
		"ng",
		"ngo",
		"nhk",
		"ni",
		"nico",
		"ninja",
		"nissan",
		"nl",
		"no",
		"np",
		"nr",
		"nra",
		"nrw",
		"ntt",
		"nu",
		"nyc",
		"nz",
		"office",
		"okinawa",
		"om",
		"omega",
		"one",
		"ong",
		"onl",
		"online",
		"ooo",
		"oracle",
		"orange",
		"org",
		"organic",
		"osaka",
		"otsuka",
		"ovh",
		"pa",
		"page",
		"panerai",
		"paris",
		"partners",
		"parts",
		"party",
		"pe",
		"pf",
		"pg",
		"ph",
		"pharmacy",
		"philips",
		"photo",
		"photography",
		"photos",
		"physio",
		"piaget",
		"pics",
		"pictet",
		"pictures",
		"pink",
		"pizza",
		"pk",
		"pl",
		"place",
		"play",
		"plumbing",
		"plus",
		"pm",
		"pn",
		"pohl",
		"poker",
		"porn",
		"post",
		"pr",
		"praxi",
		"press",
		"pro",
		"prod",
		"productions",
		"prof",
		"properties",
		"property",
		"ps",
		"pt",
		"pub",
		"pw",
		"py",
		"qa",
		"qpon",
		"quebec",
		"racing",
		"re",
		"realtor",
		"realty",
		"recipes",
		"red",
		"redstone",
		"rehab",
		"reise",
		"reisen",
		"reit",
		"ren",
		"rent",
		"rentals",
		"repair",
		"report",
		"republican",
		"rest",
		"restaurant",
		"review",
		"reviews",
		"rich",
		"ricoh",
		"rio",
		"rip",
		"ro",
		"rocks",
		"rodeo",
		"rs",
		"rsvp",
		"ru",
		"ruhr",
		"run",
		"rw",
		"ryukyu",
		"sa",
		"saarland",
		"sakura",
		"sale",
		"samsung",
		"sandvik",
		"sandvikcoromant",
		"sap",
		"sarl",
		"saxo",
		"sb",
		"sc",
		"sca",
		"scb",
		"schmidt",
		"scholarships",
		"school",
		"schule",
		"schwarz",
		"science",
		"scor",
		"scot",
		"sd",
		"se",
		"seat",
		"sener",
		"services",
		"sew",
		"sex",
		"sexy",
		"sg",
		"sh",
		"shiksha",
		"shoes",
		"show",
		"shriram",
		"si",
		"singles",
		"site",
		"sj",
		"sk",
		"ski",
		"sky",
		"skype",
		"sl",
		"sm",
		"sn",
		"sncf",
		"so",
		"soccer",
		"social",
		"software",
		"sohu",
		"solar",
		"solutions",
		"sony",
		"soy",
		"space",
		"spiegel",
		"spreadbetting",
		"sr",
		"ss",
		"st",
		"starhub",
		"statoil",
		"studio",
		"study",
		"style",
		"su",
		"sucks",
		"supplies",
		"supply",
		"support",
		"surf",
		"surgery",
		"suzuki",
		"sv",
		"swatch",
		"swiss",
		"sx",
		"sy",
		"sydney",
		"systems",
		"sz",
		"taipei",
		"tatar",
		"tattoo",
		"tax",
		"taxi",
		"tc",
		"td",
		"team",
		"tech",
		"technology",
		"tel",
		"telefonica",
		"temasek",
		"tennis",
		"tf",
		"tg",
		"th",
		"thd",
		"theater",
		"tickets",
		"tienda",
		"tips",
		"tires",
		"tirol",
		"tj",
		"tk",
		"tl",
		"tm",
		"tn",
		"to",
		"today",
		"tokyo",
		"tools",
		"top",
		"toray",
		"toshiba",
		"tours",
		"town",
		"toys",
		"tp",
		"tr",
		"trade",
		"trading",
		"training",
		"travel",
		"trust",
		"tt",
		"tui",
		"tv",
		"tw",
		"tz",
		"ua",
		"ubs",
		"ug",
		"uk",
		"um",
		"university",
		"uno",
		"uol",
		"us",
		"uy",
		"uz",
		"va",
		"vacations",
		"vc",
		"ve",
		"vegas",
		"ventures",
		"versicherung",
		"vet",
		"vg",
		"vi",
		"viajes",
		"video",
		"villas",
		"vision",
		"vista",
		"vistaprint",
		"vlaanderen",
		"vn",
		"vodka",
		"vote",
		"voting",
		"voto",
		"voyage",
		"vu",
		"wales",
		"walter",
		"wang",
		"watch",
		"webcam",
		"website",
		"wed",
		"wedding",
		"weir",
		"wf",
		"whoswho",
		"wien",
		"wiki",
		"williamhill",
		"win",
		"windows",
		"wme",
		"work",
		"works",
		"world",
		"ws",
		"wtc",
		"wtf",
		"xbox",
		"xerox",
		"xin",
		"测试",
		"परीक्षा",
		"佛山",
		"慈善",
		"集团",
		"在线",
		"한국",
		"ভারত",
		"八卦",
		"موقع",
		"বাংলা",
		"公益",
		"公司",
		"移动",
		"我爱你",
		"москва",
		"испытание",
		"қаз",
		"онлайн",
		"сайт",
		"срб",
		"бел",
		"时尚",
		"테스트",
		"淡马锡",
		"орг",
		"삼성",
		"சிங்கப்பூர்",
		"商标",
		"商店",
		"商城",
		"дети",
		"мкд",
		"טעסט",
		"工行",
		"中文网",
		"中信",
		"中国",
		"中國",
		"娱乐",
		"谷歌",
		"భారత్",
		"ලංකා",
		"測試",
		"ભારત",
		"भारत",
		"آزمایشی",
		"பரிட்சை",
		"网店",
		"संगठन",
		"餐厅",
		"网络",
		"укр",
		"香港",
		"δοκιμή",
		"飞利浦",
		"إختبار",
		"台湾",
		"台灣",
		"手机",
		"мон",
		"الجزائر",
		"عمان",
		"ایران",
		"امارات",
		"بازار",
		"پاکستان",
		"الاردن",
		"بھارت",
		"المغرب",
		"السعودية",
		"سودان",
		"عراق",
		"مليسيا",
		"澳門",
		"政府",
		"شبكة",
		"გე",
		"机构",
		"组织机构",
		"健康",
		"ไทย",
		"سورية",
		"рус",
		"рф",
		"تونس",
		"みんな",
		"グーグル",
		"ελ",
		"世界",
		"ਭਾਰਤ",
		"网址",
		"游戏",
		"vermögensberater",
		"vermögensberatung",
		"企业",
		"信息",
		"مصر",
		"قطر",
		"广东",
		"இலங்கை",
		"இந்தியா",
		"հայ",
		"新加坡",
		"فلسطين",
		"テスト",
		"政务",
		"xxx",
		"xyz",
		"yachts",
		"yandex",
		"ye",
		"yodobashi",
		"yoga",
		"yokohama",
		"youtube",
		"yt",
		"za",
		"zip",
		"zm",
		"zone",
		"zuerich",
		"zw",
	}
	FAKE_EMAIL_POOL = []byte("abcdefghijklmnopqrstuvwxyz+.-_")
)

var (
	fake_lock = sync.Mutex{}
)

func Fake(in interface{}) (err error) {
	fake_lock.Lock()
	defer fake_lock.Unlock()

	value := reflect.ValueOf(in)

	if value.Kind() != reflect.Ptr {
		err = fmt.Errorf("should pass the pointer: %T", in)
		return
	} else if !value.Elem().CanSet() {
		err = fmt.Errorf("pass %T cannot be set", in)
		return
	}

	err = fake(value.Elem())
	return
}

func fake(value reflect.Value) (err error) {
	switch kind := value.Kind(); kind {
	case reflect.Bool:
		value.SetBool(generator.Int63()%2 == 0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value.SetInt(generator.Int63())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value.SetUint(uint64(generator.Int63()))
	case reflect.Float64, reflect.Float32:
		value.SetFloat(generator.Float64())
	case reflect.String:
		size := int(generator.Int63() % FAKE_MAX_SIZE)
		data := fakeBytes(size, nil)
		value.SetString(string(data))
	case reflect.Slice, reflect.Array:
		size := value.Len()
		if size == 0 {
			// override the len to FAKE_MAX_SIZE
			size = FAKE_MAX_SIZE
		}

		for idx := 0; idx < size; idx++ {
			switch {
			case idx < value.Len():
				if err = fake(value.Index(idx)); err != nil {
					err = fmt.Errorf("cannot set #%d on %v: %v", idx, value.Type(), err)
					return
				}
			default:
				val := reflect.New(value.Type().Elem())
				if err = fake(val.Elem()); err != nil {
					err = fmt.Errorf("cannot set new instance %v: %v", val.Type(), err)
					return
				}
				value.Set(reflect.Append(value, val.Elem()))
			}
		}
		return
	case reflect.Struct:
		for idx := 0; idx < value.NumField(); idx++ {
			field := value.Field(idx)

			tags := value.Type().Field(idx).Tag
			if field.IsValid() && field.CanSet() {
				switch {
				case strings.TrimSpace(string(tags)) == FAKE_IGNORE:
				case tags.Get(FAKE_KEY) == FAKE_VALUE_NAME && field.Kind() == reflect.String:
					name := string(fakeBytes(FAKE_MAX_SIZE, FAKE_NAME_POOL))
					strings.ToTitle(name)
					field.SetString(name)
				case tags.Get(FAKE_KEY) == FAKE_VALUE_EMAIL && field.Kind() == reflect.String:
					email := string(fakeBytes(FAKE_MAX_SIZE, FAKE_EMAIL_POOL))
					email += "@" + string(fakeBytes(3, FAKE_NAME_POOL)) + "."
					email += FAKE_DOMAIN_POOL[int(generator.Int63())%len(FAKE_DOMAIN_POOL)]
					field.SetString(email)
				default:
					fake(field)
				}
			}
		}
	default:
		err = fmt.Errorf("cannot set fake for reflect.Kind: %v", kind)
		return
	}

	return
}

func fakeBytes(size int, pool []byte) (out []byte) {
	switch l := int64(len(pool)); l {
	case 0:
		for idx := 0; idx < size; idx++ {
			out = append(out, byte(generator.Int63()))
		}
	default:
		for idx := 0; idx < size; idx++ {
			out = append(out, pool[generator.Int63()%l])
		}
	}

	return
}
