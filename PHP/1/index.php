<?php

// Le dois numeros do teclado
$a = (int) trim(fgets(STDIN));
$b = (int) trim(fgets(STDIN));

// Soma dois valores inteiros
function soma(int $x, int $y): int {
    return $x + $y;
}

// Retorna o resultado da soma
$resultado = soma($a, $b);

// Exibi o resultado
echo "SOMA = $resultado\n";