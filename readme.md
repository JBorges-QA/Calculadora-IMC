# Calculadora de IMC em Go

![Versão Go](https://img.shields.io/badge/go-1.16+-00ADD8?style=for-the-badge&logo=go)
![Licença](https://img.shields.io/badge/license-MIT-green?style=for-the-badge)

Uma calculadora de Índice de Massa Corporal (IMC) de linha de comando escrita em Go que calcula e classifica seu IMC de acordo com os padrões da OMS.

## 📥 Instalação

### Pré-requisitos
- Go 1.16+ instalado

### Instalação via go install
```bash
go install github.com/seu-usuario/calculadora-imc@latest
```
Compilar manualmente
```bash
git clone https://github.com/seu-usuario/calculadora-imc.git
cd calculadora-imc
go build -o calculadora-imc
```
🚀 Como Usar
Sintaxe básica
```bash
./calculadora-imc -peso <peso_em_kg> -altura <altura_em_metros>
```
Exemplos de uso
```bash
# Calcular IMC para 70kg e 1.75m
./calculadora-imc -peso 70 -altura 1.75

# Calcular IMC para 85kg e 1.80m
./calculadora-imc -peso 85 -altura 1.80
```
## ⚙️ Opções

| Flag       | Descrição               | Unidade  | Padrão |
|------------|-------------------------|----------|--------|
| `-peso`    | Peso da pessoa          | kg       | 0      |
| `-altura`  | Altura da pessoa        | metros   | 0      |

## 📊 Classificação do IMC

| IMC          | Classificação          | Grau de Obesidade |
|--------------|------------------------|-------------------|
| < 18.5       | Magreza                | 0                 |
| 18.5 - 24.9  | Normal                 | 0                 |
| 25 - 29.9    | Sobrepeso              | 1                 |
| 30 - 39.9    | Obesidade              | 2                 |
| ≥ 40         | Obesidade Grave        | 3                 |
