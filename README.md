# Golang Course: Actividad

## Descripción
El archivo de usuarios que definimos previamente, resultó ser un registro de clientes que solicitaron una tarjeta de crédito de diversos bancos (como VISA o MASTERCARD) por lo que nuestra misión es escribir un programa que genere los números de las tarjetas a 16 dígitos de los cuales deben de seguir la siguiente nomenclatura:

- 2 dígitos indican el ID del Banco (definidos mas adelante)
- 1 dígito que indica el segmento de la tarjeta (definido mas adelante)
- 4 dígitos que indican el país de emisión. Siempre tendrá el valor 5211 el cual corresponde a México porque solo emitimos tarjetas para este país
- 9 dígitos que representan el numero de cuenta que tendremos que generar de manera aleatoria (requieren validación definida mas adelante)

### ID de Banco:
- 11: Banca Bunsan
- 22: Go Bank
- 33: Aldo Institución Financiera
- 44: Angel Banca Privada
- 55: Hugo Finantial Services
- 66: BanHumberto
- 77: Gustavo Commercial Bank

### Segmentos de tarjeta:
- 2: Tarjeta básica
- 4: Tarjeta departamental
- 6: Tarjeta clásica
- 8: Tarjeta Oro
- 0: Tarjeta Platino

###  Validación de Número de Cuenta

Considerando un numero de cuenta 345882865, lo primero a considerar es determinar las posiciones de cada uno de los dígitos. En la tabla siguiente se asigna un nombre a cada una de las posiciones (dígitos) del numero de cuenta.


position name : d9 d8 d7 d6 d5 d4 d3 d2 d1
 
account number: 3  4  5  8  8  2  8  6  5

Una vez definidos los nombres de los dígitos, el algoritmo de validación es el siguiente

 checksum calculation:

 (1*d1 + 2*d2 + 3*d3 + … + 9*d9) mod 11 = 0

En el algoritmo anterior se ordenan los nombres de posiciones (deltas) en orden ascendente y son multiplicadas por el indice que representan. El resultado del producto debe ser sumado para finalmente obtener el residuo obtenido tras la división entre 11 (módulo). Cuando el residuo o Módulo es 0 sabemos que es un numero de cuenta valido por lo que es viable asignarlo a un nuevo¡mero de tarjeta de crédito.

### Resultado
Nuestros números de cuenta deben de escribirse en un archivo en formato json.

