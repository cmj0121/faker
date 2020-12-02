package faker

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestFakerInt(t *testing.T) {
	var x int
	x_answers := []int{
		8717895732742165505,
		2259404117704393152,
		6050128673802995827,
		501233450539197794,
		3390393562759376202,
	}

	Seed(0)
	for _, ans := range x_answers {
		if err := Fake(&x); err != nil {
			t.Fatalf("cannot set faker to %T: %v", x, err)
		} else if x != ans {
			t.Fatalf("fake %v <> %v", x, ans)
		}
	}

	var y int8
	y_answers := []int8{
		-81,
		57,
		120,
		81,
	}

	rand.Seed(0)
	for _, ans := range y_answers {
		if err := Fake(&y); err != nil {
			t.Fatalf("cannot set faker to %T: %v", y, err)
		} else if y != ans {
			t.Fatalf("fake %v <> %v", y, ans)
		}
	}
}

func TestFakerFloat(t *testing.T) {
	var x float64
	x_answers := []float64{
		0.9451961492941164,
		0.24496508529377975,
		0.6559562651954052,
		0.05434383959970039,
		0.36758720663245853,
	}

	Seed(0)
	for _, ans := range x_answers {
		if err := Fake(&x); err != nil {
			t.Fatalf("cannot set faker to %T: %v", x, err)
		} else if x != ans {
			t.Fatalf("fake %v <> %v", x, ans)
		}
	}

	var y float32
	y_answers := []float32{
		0.28948045,
		0.1924386,
		0.65533215,
		0.8971697,
	}

	rand.Seed(0)
	for _, ans := range y_answers {
		if err := Fake(&y); err != nil {
			t.Fatalf("cannot set faker to %T: %v", y, err)
		} else if y != ans {
			t.Fatalf("fake %v <> %v", y, ans)
		}
	}
}

func TestFakerString(t *testing.T) {
	var x string
	x_answers := []string{
		"\xc0",
		"bJ\xaf",
		"xQN\xf8D;\xb2\xa8Y",
		"_\xc3\xccj\xf2mZ",
		" \x92o\x04k\xaaf쑥",
	}

	Seed(0)
	for _, ans := range x_answers {
		if err := Fake(&x); err != nil {
			t.Fatalf("cannot set faker to %T: %v", x, err)
		} else if x != ans {
			t.Fatalf("fake %#v <> %#v", x, ans)
		}
	}
}

func TestFakerSlice(t *testing.T) {
	var x [4]byte
	x_answers := [][4]byte{
		{0x01, 0xc0, 0x73, 0x62},
		{0x4a, 0xaf, 0x39, 0x78},
		{0x51, 0x4e, 0xf8, 0x44},
		{0x3b, 0xb2, 0xa8, 0x59},
		{0xc7, 0x5f, 0xc3, 0xcc},
	}

	Seed(0)
	for _, ans := range x_answers {
		if err := Fake(&x); err != nil {
			t.Fatalf("cannot set faker to %T: %v", x, err)
		} else if !reflect.DeepEqual(x, ans) {
			t.Fatalf("fake %#v <> %#v", x, ans)
		}
	}

	var y [7]byte
	y_answers := [][7]byte{
		{0x01, 0xc0, 0x73, 0x62, 0x4a, 0xaf, 0x39},
		{0x78, 0x51, 0x4e, 0xf8, 0x44, 0x3b, 0xb2},
		{0xa8, 0x59, 0xc7, 0x5f, 0xc3, 0xcc, 0x6a},
		{0xf2, 0x6d, 0x5a, 0xaa, 0x20, 0x92, 0x6f},
		{0x04, 0x6b, 0xaa, 0x66, 0xec, 0x91, 0xa5},
	}

	Seed(0)
	for _, ans := range y_answers {
		if err := Fake(&y); err != nil {
			t.Fatalf("cannot set faker to %T: %v", y, err)
		} else if !reflect.DeepEqual(y, ans) {
			t.Fatalf("fake %#v <> %#v", y, ans)
		}
	}

	var z []byte
	z_answers := [][]byte{
		{0x01, 0xc0, 0x73, 0x62, 0x4a, 0xaf, 0x39, 0x78, 0x51, 0x4e, 0xf8, 0x44, 0x3b, 0xb2, 0xa8, 0x59},
		{0xc7, 0x5f, 0xc3, 0xcc, 0x6a, 0xf2, 0x6d, 0x5a, 0xaa, 0x20, 0x92, 0x6f, 0x04, 0x6b, 0xaa, 0x66},
		{0xec, 0x91, 0xa5, 0xa2, 0x79, 0x43, 0x23, 0xc2, 0xda, 0x40, 0x5a, 0xfe, 0xd6, 0x1e, 0xc1, 0x5c},
		{0xf7, 0x93, 0xed, 0x10, 0x3a, 0x57, 0xc0, 0x2c, 0x86, 0x9a, 0xf0, 0x69, 0xb7, 0xb5, 0x68, 0x37},
		{0xa2, 0x52, 0x7a, 0x30, 0xb2, 0x77, 0xe5, 0x5a, 0x4d, 0xc6, 0xf3, 0x35, 0xf1, 0xfb, 0xe9, 0x9f},
	}

	Seed(0)
	for _, ans := range z_answers {
		if err := Fake(&z); err != nil {
			t.Fatalf("cannot set faker to %T: %v", z, err)
		} else if !reflect.DeepEqual(z, ans) {
			t.Fatalf("fake %#v <> %#v", z, ans)
		}
	}
}

func BenchmarkFakeInt(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x int

		for pb.Next() {
			// generate the fake int
			Fake(&x)
		}
	})
}

func BenchmarkFakeFloat(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x float64

		for pb.Next() {
			// generate the fake int
			Fake(&x)
		}
	})
}

func BenchmarkFakeString(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x string

		for pb.Next() {
			// generate the fake int
			Fake(&x)
		}
	})
}

func BenchmarkFakeSlice32(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x [32]byte

		for pb.Next() {
			// generate the fake int
			Fake(&x)
		}
	})
}

func BenchmarkFakeSlice64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x [64]byte

		for pb.Next() {
			// generate the fake int
			Fake(&x)
		}
	})
}

func BenchmarkFakeSlice(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x []byte

		for pb.Next() {
			// generate the fake int
			Fake(&x)
		}
	})
}
