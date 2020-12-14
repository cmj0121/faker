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
	FAKE_TAG_FLAG     = "fake"
	FAKE_VALUE_NAME   = "name"
	FAKE_VALUE_DOMAIN = "domain"
	FAKE_VALUE_EMAIL  = "email"
	FAKE_VALUE_LOWER  = "lower"
	FAKE_VALUE_UPPER  = "upper"
	FAKE_VALUE_DIGIT  = "digit"
)

var (
	FAKE_NAME_LISTS = []string{
		"john",
		"cindy",
		"きんだいち",
	}
	FAKE_DOMAIN_LISTS = []string{
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

	FAKE_EMAIL_LISTS = []string{
		"%[1]v@%[2]v",
		"%[1]v+%[3]d@%[2]v",
		"%[1]v+%[3]d@%[4]d.%[2]v",
	}
)
