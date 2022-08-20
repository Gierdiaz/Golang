// TESTE DE UNIDADE
package enderecos

import "testing"

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoEndereco(t *testing.T) {
	
	cenariosTestes := []cenarioTeste {
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Estrada do cabucu", "Estrada"},
		{"Praça das Rosas", "Tipo inválido"},
		{" ", "Tipo inválido"},
		{"AVENIDA PREBOUÇA", "Avenida"},
	}

	for _, cenario := range cenariosTestes {
		retornoRecebido := TipoEndereco(cenario.enderecoInserido) 
			if retornoRecebido != cenario.retornoEsperado {
				t.Errorf("O Tipo recebido %s é diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
			}
		
	}

}