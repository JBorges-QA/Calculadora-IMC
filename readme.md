# Conversor de Temperatura CLI em Go

![Versão Go](https://img.shields.io/badge/go-1.16+-00ADD8?style=for-the-badge&logo=go)
![Licença](https://img.shields.io/badge/license-MIT-green?style=for-the-badge)

Um conversor de temperaturas de linha de comando escrito em Go que realiza conversões precisas entre Celsius (°C), Fahrenheit (°F) e Kelvin (K).

## 📥 Instalação

### Pré-requisitos
- Go 1.16+ instalado

### Instalação via go install
```bash
go install github.com/seu-usuario/conversor-temperatura@latest
Compilar manualmente
bash
git clone https://github.com/seu-usuario/conversor-temperatura.git
cd conversor-temperatura
go build -o conversor-temp
🚀 Como Usar
Sintaxe básica
bash
./conversor-temp -valor <temperatura> -de <unidade_origem> -para <unidade_destino>
Exemplos de uso
bash
# Converter 100°C para Fahrenheit
./conversor-temp -valor 100 -de C -para F

# Converter 32°F para Kelvin
./conversor-temp -valor 32 -de F -para K

# Converter 273.15K para Celsius
./conversor-temp -valor 273.15 -de K -para C
⚙️ Opções
Flag	Descrição	Valores Válidos	Padrão
-valor	Valor da temperatura	Número real	0
-de	Unidade de temperatura original	C, F, K	C
-para	Unidade de temperatura alvo	C, F, K	F
🔄 Conversões Suportadas
De \ Para	Celsius (°C)	Fahrenheit (°F)	Kelvin (K)
Celsius	—	✅	✅
Fahrenheit	✅	—	✅
Kelvin	✅	✅	—
