
# Weather CEP API (Go + Gin)

Este projeto é um lab da pós-graduação Full Cycle onde foi desenvolvido um microserviço simples em Go para buscar a temperatura de uma cidade com base em um CEP brasileiro.

A aplicação consulta o CEP via [ViaCEP](https://viacep.com.br) e retorna a temperatura atual consultando a [WeatherAPI](https://www.weatherapi.com/), em três unidades: Celsius, Fahrenheit e Kelvin.

---

## Tecnologias Utilizadas

- Go 1.24.2
- Gin Gonic (framework web)
- WeatherAPI (clima)
- ViaCEP (localização)
- Docker
- Docker Compose

---

## 📦 Como rodar

### 🔧 Pré-requisitos

- Docker instalado
- Git instalado (opcional)

---

### Rodar com Docker Compose (recomendado)

```bash
git clone https://github.com/fjgmelloni/weather-cep.git
cd weather-cep

echo "WEATHER_API_KEY=aeca72e73a1f4acfa7f163903251104" > .env

docker compose up --build
```

---

### Rodar somente com Docker (sem Compose)

```bash
docker build -t weather-cep .
docker run --env-file .env -p 8080:8080 weather-cep
```

> Alternativamente, passe a chave diretamente:

```bash
docker run -e WEATHER_API_KEY=aeca72e73a1f4acfa7f163903251104 -p 8080:8080 weather-cep
```

---

## Como testar a API

### CEP válido

```bash
curl http://localhost:8080/weather?cep=01001000
```

**Resposta:**

```json
{
  "temp_C": 27.5,
  "temp_F": 81.5,
  "temp_K": 300.5
}
```

---

### CEP com formato inválido (ex: menos de 8 dígitos)

```bash
curl http://localhost:8080/weather?cep=123
```

**HTTP 422 –** `invalid zipcode`

---

### CEP não encontrado

```bash
curl http://localhost:8080/weather?cep=00000000
```

**HTTP 404 –** `can not find zipcode`

---

## Regras de negócio atendidas

- [x] Recebe um CEP com exatamente 8 dígitos
- [x] Valida formato do CEP
- [x] Busca a cidade via ViaCEP
- [x] Consulta temperatura via WeatherAPI
- [x] Retorna as temperaturas em:
  - Celsius
  - Fahrenheit
  - Kelvin
- [x] Retorna erros apropriados com HTTP 422 e 404
- [x] Pode ser executado localmente, via Docker e via Docker Compose

---
