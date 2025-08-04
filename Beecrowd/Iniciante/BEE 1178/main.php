<?php
 
// Lê o valor da entrada
$valor = (float) trim(fgets(STDIN));

// Inicia o vetor
$N = array();

// Coloca o valor na primeira posição
$N[0] = $valor;

// Preenche o restante com a metade do anterior
for ($i = 1; $i < 100; $i++) {
    $N[$i] = $N[$i - 1] / 2;
}

for ($i = 0; $i < 100; $i++) {
    printf("N[%d] = %.4f\n", $i, $N[$i]);
}

