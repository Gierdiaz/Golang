// TESTE DE UNIDADE
package enderecos

import "testing"

func TestTipoEndereco(t *testing.T) {
	// tem que ser exatamente como está. Com a primeira e a segunda letra LOWER
	enderecoParaTeste := "Rua ABC"

	tipoEnderecoEsperado := "Avenida"

	tipoEnderecoRecebido := tipoEndereco(enderecoParaTeste)

	if tipoEnderecoRecebido != tipoEnderecoEsperado {
		t.Errorf("O Tipo recebido é diferente do esperado. Esperava %s e recebeu %s", tipoEnderecoEsperado, tipoEnderecoRecebido)
	}

}