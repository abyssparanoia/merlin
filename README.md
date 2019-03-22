# merilin

## devOps

### prepare env

```bash
> cp env.example.json env.json
```

### install

```bash
> dep ensure
> make init
```

### local dev

- migration

```bash
> docker-compose build
```

- up mysql

```bash
> docker-compose up -d
```

- down mysql

```bash
> docker-compose down
```

- start server

```bash
> make run app=api
```
