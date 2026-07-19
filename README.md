# Proyecto_2B — API REST E-commerce (Go + Gin)

API RESTful de e-commerce desarrollada en **Go** con **Gin**, **GORM** y **PostgreSQL**, como parte de la tarea integradora de Aplicaciones Web (Grupo 8 — Escuela Politécnica Nacional). Implementa el mismo dominio, reglas de negocio y endpoints que el backend base de referencia (Spring Boot), adaptados al stack Go.

## Descripción

Sistema de e-commerce con gestión de usuarios, productos y recibos de compra (órdenes). El backend calcula el total de cada compra a partir del precio real almacenado en base de datos, valida el stock disponible y lo descuenta automáticamente al confirmar un recibo.

## Arquitectura

Arquitectura por capas:

```
main.go            # bootstrap: conexión DB, migraciones, wiring de dependencias, router
config/            # conexión a PostgreSQL (GORM)
models/            # entidades GORM (User, Product, Receipt, ReceiptItem)
dto/                # objetos de entrada/salida (requests, responses, error estándar)
repository/         # acceso a datos (GORM) por entidad
services/            # lógica de negocio y reglas de dominio
controllers/         # handlers HTTP (Gin) + documentación Swagger (anotaciones godoc)
middleware/          # autenticación JWT (AuthMiddleware)
utils/               # hashing de contraseñas (BCrypt) y generación/validación de JWT
exceptions/          # errores de dominio centralizados (sentinel errors)
routes/              # definición de rutas y agrupación pública/protegida
docs/                # especificación Swagger/OpenAPI generada (swag)
```

Flujo de una petición: `router → middleware (JWT si aplica) → controller (valida DTO) → service (regla de negocio) → repository (GORM/Postgres) → respuesta JSON`.

## Modelo de datos

- **User**: `userId, firstName, lastName, email (único), password (hash BCrypt, nunca se serializa), address, phoneNumber`
- **Product**: `productId, name, price (decimal), description, amount (stock), imageUrl`
- **Receipt**: `receiptId, userId (FK), total (decimal, calculado), amountOfItems, createdAt`
- **ReceiptItem**: `receiptItemId, receiptId (FK), productId (FK), quantity, unitPrice, subtotal`

Las tablas se crean con `AutoMigrate` de GORM al iniciar la aplicación (`main.go`); las relaciones (`FOREIGN KEY`) se agregan explícitamente vía SQL en el arranque.

## Reglas de negocio implementadas

- El cliente **nunca** envía el total del recibo; se calcula en `services/receipt_services.go` sumando `precio_unitario_bd × cantidad` de cada ítem.
- Al crear un recibo se valida stock disponible por producto (`product.Amount < quantity` → error `stock insuficiente`, HTTP 400).
- Al confirmar el recibo, el stock del producto se descuenta automáticamente dentro de una transacción de base de datos (todo o nada).
- Las contraseñas se cifran con **BCrypt** (`utils/password.go`) y jamás se devuelven en las respuestas HTTP (`json:"-"` en el modelo `User`).
- Los montos monetarios usan `decimal.Decimal` (`shopspring/decimal`) mapeado a `numeric(10,2)` en PostgreSQL, evitando errores de precisión de punto flotante.
- Manejo de errores centralizado: los servicios devuelven errores de dominio (`exceptions/errors.go`) que `controllers/error_handler.go` traduce a respuestas JSON consistentes (`dto.APIError`) con timestamp, status, error, mensaje y path.

## Seguridad

- Login devuelve un **JWT** (`utils/jwt.go`) que debe enviarse como `Authorization: Bearer <token>`.
- Rutas protegidas con `middleware.AuthMiddleware()`:
  - `GET/PUT/DELETE /api/users/{id}`
  - `POST/PUT/DELETE /api/products` (lectura de productos es pública)
  - Todas las rutas de `/api/receipts`
- `POST /api/users/register` y `POST /api/users/login` son públicas.

## Requisitos previos

- [Go 1.26+](https://go.dev/dl/) (o usar Docker, ver más abajo — no requiere Go instalado)
- PostgreSQL 13+ (local, o vía Docker)

## Variables de entorno

Copiar `.env.example` a `.env` y ajustar los valores:

| Variable | Descripción | Ejemplo |
|---|---|---|
| `DB_HOST` | Host de PostgreSQL | `localhost` |
| `DB_USER` | Usuario de PostgreSQL | `postgres` |
| `DB_PASSWORD` | Contraseña de PostgreSQL | `postgres` |
| `DB_NAME` | Nombre de la base de datos | `proyecto2b` |
| `DB_PORT` | Puerto de PostgreSQL | `5432` |
| `JWT_SECRET` | Clave para firmar los tokens JWT | cualquier cadena secreta larga |

## Ejecución

### Opción A — Docker (recomendado, no requiere instalar Go ni Postgres)

```bash
cp .env.example .env
docker compose up --build
```

Esto levanta un contenedor de PostgreSQL y otro con la API en `http://localhost:8080`.

### Opción B — Local con Go instalado

```bash
cp .env.example .env
# Editar .env con los datos de tu PostgreSQL local
go mod download
go run .
```

El servidor queda escuchando en `http://localhost:8080`.

## Documentación de la API

- Swagger UI: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- Especificación OpenAPI: `docs/swagger.json` / `docs/swagger.yaml`

Para regenerar la documentación tras cambiar las anotaciones godoc:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

## Pruebas

Colección de Postman incluida en [`postman/Proyecto_2B.postman_collection.json`](postman/Proyecto_2B.postman_collection.json), con variables `{{baseUrl}}` y `{{token}}` (este último se guarda automáticamente tras el login).

Flujo de prueba sugerido: `Register → Login (guarda token) → Create Product → Create Receipt → Get Receipt`.

## Endpoints

### Users

| Método | Ruta | Auth | Descripción |
|---|---|---|---|
| POST | `/api/users/register` | No | Registro de usuario |
| POST | `/api/users/login` | No | Login, devuelve JWT |
| GET | `/api/users/{id}` | Sí | Consultar usuario |
| PUT | `/api/users/{id}` | Sí | Actualizar usuario |
| DELETE | `/api/users/{id}` | Sí | Eliminar usuario |

### Products

| Método | Ruta | Auth | Descripción |
|---|---|---|---|
| POST | `/api/products` | Sí | Crear producto |
| GET | `/api/products` | No | Listar productos |
| GET | `/api/products/{id}` | No | Buscar producto por id |
| PUT | `/api/products/{id}` | Sí | Actualizar producto |
| DELETE | `/api/products/{id}` | Sí | Eliminar producto |

### Receipts

| Método | Ruta | Auth | Descripción |
|---|---|---|---|
| POST | `/api/receipts` | Sí | Crear recibo (calcula total y descuenta stock) |
| GET | `/api/receipts` | Sí | Listar todos los recibos |
| GET | `/api/receipts/{id}` | Sí | Buscar recibo por id |
| GET | `/api/receipts/user/{userId}` | Sí | Listar recibos de un usuario |
| DELETE | `/api/receipts/{id}` | Sí | Eliminar recibo |

## Stack técnico

- **Lenguaje**: Go 1.26
- **Framework HTTP**: Gin
- **ORM**: GORM
- **Base de datos**: PostgreSQL
- **Autenticación**: JWT (`golang-jwt/jwt`)
- **Hashing de contraseñas**: BCrypt (`golang.org/x/crypto/bcrypt`)
- **Decimales**: `shopspring/decimal`
- **Documentación**: Swagger/OpenAPI (`swaggo/swag`, `swaggo/gin-swagger`)
