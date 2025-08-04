<?php

// Verificar se a string contém uma palavra específica

$entrada = (string) trim(fgets(STDIN));

$saida = str_contains(strtolower($entrada), 'nome');

echo "Saída: $saida";