<?php

// Primeira letra maiúscula

$entrada = (string) trim(fgets(STDIN));

$saida = ucfirst($entrada);

echo "Saída: $saida";