# Assessment MAy 2024

A brief description of what the project is about.

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yumekun/Assessment-may-2024
    ```

2. Change to the project directory:
    ```bash
    cd Assessment-may-2024
    ```

3. Create a `.env` file in the root directory of your project with the following content:
    ```plaintext
    # Environment variables
    POSTGRES_PASSWORD=changeme
    POSTGRES_USER=postgres
    POSTGRES_DB=postgres
    ```

4. Create a `config.env` file in each service directory with the following content:
    ```plaintext
    # Service configuration
    POSTGRES_DRIVER=postgres
    POSTGRES_URL=postgresql://postgres:changeme@postgres:5432/postgres?sslmode=disable

    REDIS_SERVICE_ADDRESS=redis:6379
    REDIS_PASSWORD=

    REDIS_MUTASI_REQUEST_STREAM=mutasi_req
    ```

5. run docker compose
```bash
   docker compose up
    ```
