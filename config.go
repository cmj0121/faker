package faker

// program meta
const (
	// program name
	PROJ_NAME = "faker"
	// version info
	MAJOR = 0
	MINOR = 2
	MACRO = 0
)

const (
	// the default slice max size
	FAKE_MAX_SLICE_LEN = 16
	// the ignore tag when fake
	FAKE_TAG_IGNORE = "-"

	// the customized tag
	// the fixed-length slice or string
	FAKE_TAG_SLICE_SIZE = "fake_size"
	// the fixed-choice string
	FAKE_TAG_STR_ENUM = "fake_enum"
	FAKE_VALUE_NAME   = "name"
	FAKE_VALUE_DOMAIN = "domain"
)

var (
	FAKE_NAME_POOL = []string{
		"john",
		"cindy",
		"きんだいち",
	}
	FAKE_DOMAIN_POOL = []string{
		"com",
		"com.tw",
		"example",
		"vermögensberatung",
		"இந்தியா",
		"فلسطين",
		"テスト",
		// https://jasontucker.blog/8945/what-is-the-longest-tld-you-can-get-for-a-domain-name
		"XN--VERMGENSBERATUNG-PWB",
	}
)
