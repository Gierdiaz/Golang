<?php

enum Gesto {
    case PEDRA;
    case PAPEL;
    case TESOURA;
}

enum Resultado {
    case GANHA;
    case PERDE;
    case EMPATA;
}

function padraPapelTesoura(string $entrada): ?Gesto {
    return match (strtolower(trim($entrada))) {
        'pedra' => Gesto::PEDRA,
        'papel' => Gesto::PAPEL,
        'tesoura' => Gesto::TESOURA,
        default => null,
    };
}

function maos(Gesto $jogadorA, Gesto $jogadorB): Resultado {
    if ($jogadorA === $jogadorB ) {
        return Resultado::EMPATA;
    }

    return match ([$jogadorA, $jogadorB]) {
        [Gesto::PEDRA, Gesto::TESOURA] => Resultado::GANHA,
        [Gesto::TESOURA, Gesto::PAPEL]  => Resultado::GANHA,
        [Gesto::PAPEL, Gesto::PEDRA] => Resultado::GANHA,
        default => Resultado::PERDE,
    };
}

$entradaA = (string) trim(fgets(STDIN));
$entradaB = (string) trim(fgets(STDIN));

$jogadorA = padraPapelTesoura($entradaA);
$jogadorB = padraPapelTesoura($entradaB);

if (is_null($jogadorA) || is_null($jogadorB)) {
    echo 'Entrada Inválida!\n';
    exit;
}

$resultado = maos($jogadorA, $jogadorB);

echo match ($resultado) {
    Resultado::EMPATA => "Empate!\n",
    Resultado::GANHA  => "Jogador A venceu! Jogador B perdeu!\n",
    Resultado::PERDE  => "Jogador B venceu! Jogador A perdeu!\n"
};