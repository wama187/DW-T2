# S10-P1

Este proyecto es una aplicación web con un frontend en Vue.js y un backend en Go.

## Prerrequisitos

*   Docker
*   Docker Compose

## Ejecutar la aplicación

1.  **Clonar el repositorio:**

    ```bash
    git clone https://github.com/Crow-star/S10-P1.git
    cd S10-P1
    ```

2.  **Crear un archivo `.env` en el directorio `client` con las siguientes variables:**

    ```bash
    GOOGLE_CLIENT=YOUR_GOOGLE_CLIENT_ID
    GOOGLE_SECRET=YOUR_GOOGLE_SECRET
    GOOGLE_REDIRECT_URL=http://localhost:8080/auth/google/callback
    ```

3.  **Ejecutar la aplicación usando Docker Compose:**

    ```bash
    docker-compose up -d
    ```

    Esto construirá las imágenes de Docker y ejecutará los contenedores para el frontend y el backend.

4.  **Acceder a la aplicación:**

    *   Frontend: [http://localhost:8080](http://localhost:8080)
    *   Backend: [http://localhost:8000](http://localhost:8000)
