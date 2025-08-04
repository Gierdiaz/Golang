<?php

// Quantos caracteres tem a string?

// trim() -> remove espacos e quabras de linha no inicio e no final de uma string
// fgets(STDIN) -> le a entrada com possivel quebra de linha no final
// strlen() -> cotna todos os caracteres, inclusive espacos nno meio da string

$valor = (string) trim(fgets(STDIN));

$total = strlen($valor);

echo "Total de caracteres: $total\n";