package faker

import (
	"math/rand"
	"testing"
)

/* ---- general type ---- */
func TestFakerBool(t *testing.T) {
	var x bool
	x_answers := []bool{
		false,
		true,
		false,
		true,
		true,
	}

	Seed(0)
	for _, ans := range x_answers {
		if err := Fake(&x); err != nil {
			t.Fatalf("cannot set faker to %T: %v", x, err)
		} else if x != ans {
			t.Fatalf("fake %v <> %v", x, ans)
		}
	}
}

func TestFakerByte(t *testing.T) {
	var x byte
	x_answers := []byte{
		1,
		192,
		115,
		98,
		74,
	}

	Seed(0)
	for _, ans := range x_answers {
		if err := Fake(&x); err != nil {
			t.Fatalf("cannot set faker to %T: %v", x, err)
		} else if x != ans {
			t.Fatalf("fake %v <> %v", x, ans)
		}
	}
}

func TestFakerRune(t *testing.T) {
	var x rune
	x_answers := []rune{
		-1023568895,
		-2932288,
		1332660339,
		-387013278,
		1963666762,
	}

	Seed(0)
	for _, ans := range x_answers {
		if err := Fake(&x); err != nil {
			t.Fatalf("cannot set faker to %T: %v", x, err)
		} else if x != ans {
			t.Fatalf("fake %v <> %v", x, ans)
		}
	}
}

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

func TestFakerComplex(t *testing.T) {
	var x complex128
	x_answers := []complex128{
		0.9451961492941164 + 0.24496508529377975i,
		0.6559562651954052 + 0.05434383959970039i,
		0.36758720663245853 + 0.2894804331565928i,
		0.19243860967493215 + 0.6553321508148324i,
		0.897169713149801 + 0.16735444255905835i,
	}

	Seed(0)
	for _, ans := range x_answers {
		if err := Fake(&x); err != nil {
			t.Fatalf("cannot set faker to %T: %v", x, err)
		} else if x != ans {
			t.Fatalf("fake %v <> %v", x, ans)
		}
	}

	var y complex64
	y_answers := []complex64{
		0.94519615 + 0.24496509i,
		0.65595627 + 0.05434384i,
		0.3675872 + 0.28948045i,
		0.1924386 + 0.65533215i,
		0.8971697 + 0.16735445i,
	}

	Seed(0)
	for _, ans := range y_answers {
		if err := Fake(&y); err != nil {
			t.Fatalf("cannot set faker to %T: %v", y, err)
		} else if y != ans {
			t.Fatalf("fake %v <> %v", y, ans)
		}
	}
}

/* ---- benchmark ---- */
func BenchmarkFakeBool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x byte

		for pb.Next() {
			// generate the fake int
			MustFake(&x)
		}
	})
}

func BenchmarkFakeByte(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x int

		for pb.Next() {
			// generate the fake int
			MustFake(&x)
		}
	})
}

func BenchmarkFakeRune(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x rune

		for pb.Next() {
			// generate the fake int
			MustFake(&x)
		}
	})
}

func BenchmarkFakeInt(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x int

		for pb.Next() {
			// generate the fake int
			MustFake(&x)
		}
	})
}

func BenchmarkFakeFloat(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x float64

		for pb.Next() {
			// generate the fake int
			MustFake(&x)
		}
	})
}

func BenchmarkFakeComplex(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x complex128

		for pb.Next() {
			// generate the fake int
			MustFake(&x)
		}
	})
}

func BenchmarkCryptoRandomFakeInt(b *testing.B) {
	SetGenerator(CryptoRandom{})
	b.RunParallel(func(pb *testing.PB) {
		var x int

		for pb.Next() {
			// generate the fake int
			MustFake(&x)
		}
	})
}
