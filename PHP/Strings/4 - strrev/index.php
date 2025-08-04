<?php

// Inverter a string

$entrada = (string) trim(fgets(STDIN));

$saida = strrev($entrada);

echo "Saída: $saida";
