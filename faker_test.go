package faker

import (
	"math/rand"
	"testing"
)

func TestFakerInt(t *testing.T) {
	x := 10

	Seed(0)
	if err := Fake(&x); err != nil {
		t.Fatalf("fake %T: %v", x, err)
	} else if x != 8717895732742165505 {
		t.Fatalf("fake int: %v <> 8717895732742165505", x)
	}

	y := int8(10)

	Seed(0)
	if err := Fake(&y); err != nil {
		t.Fatalf("fake %T: %v", y, err)
	} else if y != 1 {
		t.Fatalf("fake int: %v <> 1", y)
	}
}

func TestFakerFloat(t *testing.T) {
	x := 1.2

	Seed(0)
	if err := Fake(&x); err != nil {
		t.Fatalf("fake %T: %v", x, err)
	} else if x != 0.9451961492941164 {
		t.Fatalf("fake int: %v <> 0.9451961492941164", x)
	}

	y := float32(1.2)

	Seed(0)
	if err := Fake(&y); err != nil {
		t.Fatalf("fake %T: %v", y, err)
	} else if y != 0.9451961492941164 {
		t.Fatalf("fake int: %v <> 0.9451961492941164", y)
	}
}

func TestFakerIntRepeat(t *testing.T) {
	x := 10
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

	y := int8(10)
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
