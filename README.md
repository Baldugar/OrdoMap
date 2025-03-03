# OrdoMap

Crear una plataforma integral para **worldbuilders** y directores de juego, que facilite la construcción de mundos ficticios con herramientas avanzadas de gestión de lore, creación de personajes, mapas interactivos, relaciones entre entidades y generación de contenido con inteligencia artificial.

# Iniciar Proyecto

-   Necesitas tener ArangoDB instalado o correrlo en un contenedor de Docker.
-   Necesitas tener NodeJS instalado.
-   Necesitas tener Go instalado.

## Iniciar Servidor

Hace falta tener un archivo `settings.development.json` en la carpeta `server` con la siguiente estructura:

```json
{
    "allowCrossOrigin": true,
    "logging": {
        "logLevel": "debug"
    },
    "domain": "localhost",
    "httpListen": "8080",
    "graphQLPlayground": true,
    "arangoDB": {
        "addr": "localhost",
        "port": "XXXX",
        "name": "OrdoMap",
        "user": "XXXX",
        "password": "XXXX"
    }
}
```

Donde `XXXX` es el puerto, usuario y contraseña de la base de datos.

Ejecutar el servidor:

```bash
cd server
./run.sh
```

El script `run.sh` compila el servidor y lo ejecuta.

El servidor se encarga de crear la base de datos y las colecciones necesarias. Así como de importar los datos iniciales. La primera vez que se ejecute el servidor, tardará un poco más de lo normal.

## Iniciar Cliente

```bash
cd client
npm install
npm run dev
```

El cliente se ejecuta en `http://localhost:5173/`. Puedes cambiar el puerto en el archivo `client/vue.config.js`.
