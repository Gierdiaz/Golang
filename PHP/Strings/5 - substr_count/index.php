<?php

// Quantas vezes aparece a letra "a"

// substr_count() -> Conta o número de ocorrências de uma substring

$entrada = (string) trim(fgets(STDIN));

$saida = substr_count(strtolower($entrada), 'nome');

echo "A palavra aparece $saida vezes.\n";


