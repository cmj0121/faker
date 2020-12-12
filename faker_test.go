package faker

import (
	"math/rand"
	"reflect"
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

func TestFakerArray(t *testing.T) {
	var x [3]int8
	x_answers := [][3]int8{
		{1, -64, 115},
		{98, 74, -81},
		{57, 120, 81},
		{78, -8, 68},
		{59, -78, -88},
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

func TestFakerSlice(t *testing.T) {
	x := []int8{}
	x_answers := [][]int8{
		{-64},
		{98, 74, -81},
		{120, 81, 78, -8, 68, 59, -78, -88, 89},
		{95, -61, -52, 106, -14, 109, 90, -88, 89},
		{32, -110, 111, 4, 107, -86, 102, -20, -111, -91},
		{121, 67, 111, 4, 107, -86, 102, -20, -111, -91},
	}

	Seed(0)
	for _, ans := range x_answers {
		if err := Fake(&x); err != nil {
			t.Fatalf("cannot set faker to %T: %v", x, err)
		} else if !reflect.DeepEqual(x, ans) {
			t.Fatalf("fake %#v <> %#v", x, ans)
		}
	}

	y := []float32{}
	y_answers := [][]float32{
		{0.24496509},
		{0.05434384, 0.3675872, 0.28948045},
		{0.65533215, 0.8971697, 0.16735445, 0.28858566, 0.9026048, 0.84978026, 0.2730468, 0.6090802, 0.253656},
		{0.017480763, 0.78707397, 0.7993937, 0.35640854, 0.42619205, 0.51024234, 0.2404319, 0.6090802, 0.253656},
		{0.69307005, 0.4018979, 0.2848241, 0.6833966, 0.43753722, 0.104014836, 0.3159685, 0.1512936, 0.7313419, 0.31416726},
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

func BenchmarkFakeArray(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var x [16]int

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
