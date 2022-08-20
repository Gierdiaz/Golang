package formas

import ("math"
		"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T){
		ret := Retangulo{10,12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T){
		cir := 	Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := cir.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}
	})
}