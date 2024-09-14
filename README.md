# Dark Mode CSS Serverless Function

Este proyecto proporciona una función serverless escrita en Go que genera CSS dinámico para modos oscuro y claro. La función está desplegada en Vercel y responde a solicitudes HTTP con CSS basado en el parámetro `theme` de la consulta.

## Estructura del Proyecto

- `api/theme-css.go` - Contiene el código Go para la función serverless.
- `vercel.json` - Configuración para desplegar la función en Vercel.

## Configuración

### `theme-css.go`

Esta función maneja solicitudes HTTP y devuelve CSS dinámico. El tema predeterminado es el modo oscuro, pero se puede especificar `dark` o `light` a través del parámetro `theme` en la consulta.

### `vercel.json`

Configura el entorno de ejecución para Vercel y define la ruta a la función serverless.

```json
{
  "functions": {
    "api/theme-css.go": {
      "runtime": "vercel-go@1.10.0"
    }
  }
}
