# core

# Setup

1. Init submodules

```bash
git submodule update --init --recursive --remote
```

2. Install dependency

```bash
make setup
```

3. Run sub service
```bash
docker compose up -d
```

4. Setup credentials
```bash
cp .env.example .env
```

5. Run service
```bash
make run
```