# Party games

Запускай теперь через docker-compouse.yaml

## Back

1. Build image
    ```bash
    docker build -t app .
    ```
2. Run container
    ```bash
   docker run -p 8080:8080 app
   ```

## Front

1. Собери образ:

  ```bash
    docker build -t my-react-app .
   ```

2. Запусти контейнер:

  ```bash
    docker run -d -p 3000:80 my-react-app
```

Открой браузер:
http://localhost:3000

## Documentation

Format swagger comments:

```bash
swag fmt
```

Generate documentation:
```bash
swag init -d ./cmd,./
```

You can read the API documentation at http://localhost:8080/swagger/index.html