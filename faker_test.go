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
		"bJ\xaf9xQN\xf8D;\xb2\xa8Y\xc7_\xc3\xccj\xf2",
		"Z\xaa \x92o\x04k\xaaf쑥\xa2",
		"C#\xc2\xda@Z\xfe\xd6\x1e\xc1\\\xf7\x93\xed\x10:W\xc0,\x86\x9a\xf0i\xb7\xb5",
		"7\xa2Rz0\xb2w\xe5",
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
