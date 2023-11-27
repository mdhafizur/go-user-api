# Running Docker Compose for a Specific Service

To start your Docker Compose setup for a specific service, follow these steps:

1. Open a terminal.

2. Run the following command to start Docker Compose with a specific service, rebuild the images, and force recreation:

```bash
docker compose -f compose.yml up -d --build [service-name] --force-recreate
```

```bash
docker compose -f compose.yml up -d --build go-api --force-recreate
```

```bash
docker compose -f compose.yml up -d --build mongodb --force-recreate
```
