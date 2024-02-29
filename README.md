# Conversor de Escalas Termométricas em Go

Este é um programa simples em Go que permite converter temperaturas entre Celsius, Fahrenheit e Kelvin.

## Funcionalidades

O programa oferece as seguintes funcionalidades:

- Conversão de Celsius para Fahrenheit
- Conversão de Fahrenheit para Celsius
- Conversão de Celsius para Kelvin
- Conversão de Kelvin para Celsius
- Conversão de Fahrenheit para Kelvin
- Conversão de Kelvin para Fahrenheit

## Como usar

1. Certifique-se de ter o Go instalado em seu sistema. Você pode baixá-lo em [golang.org](https://golang.org/dl/).

2. Clone ou baixe este repositório para o seu computador.

3. Abra o terminal ou prompt de comando e navegue até o diretório onde o arquivo `converter_temperatura.go` está localizado.

4. Execute o seguinte comando para compilar e executar o programa:
```bash
go run converter_temperatura.go
```
5. O programa será iniciado e solicitará que você insira uma temperatura e escolha a escala atual da temperatura (C, F ou K).

6. O programa calculará e exibirá a temperatura nas outras escalas.

## Explicação do Código

O código consiste em:

- Funções separadas para cada tipo de conversão (Celsius para Fahrenheit, Fahrenheit para Celsius, etc.).

- O uso de um loop `switch` para determinar qual conversão realizar com base na entrada do usuário.

- A entrada do usuário é lida usando a função `Scan` do pacote `fmt`.

- As conversões são realizadas usando as fórmulas comuns para cada tipo de escala termométrica.
