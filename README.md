# Rate Limiter API

A simple HTTP API with rate limiting built on top of Redis. I was messing around with rate limiting patterns and decided to implement a fixed window rate limiter in Go.

## What's this?

Basic API server that limits how many requests an IP address can make within a time window. Uses Redis to track request counts, so it'll work across multiple instances if you scale horizontally (though I haven't really tested that much yet).

## How it works

Pretty straightforward - uses the fixed window algorithm:
- Each IP gets tracked in Redis with a key like `rate:127.0.0.1`
- Counter increments on each request
- Window resets after the configured time period
- If you hit the limit, you get blocked until the window resets

Yeah, I know fixed window has that burst problem at window boundaries, but it's simple and good enough for most cases.

## Setup

You'll need Go 1.25+ and Docker (for Redis).

```bash
# Clone it
git clone <your-repo-url>
cd RateLimitor-with-go

# Install dependencies
go mod download

# Start Redis
docker-compose up -d
```

The `.env` file has some defaults:
```
PORT=":8081"
REDIS_ADDR=localhost:6379
REDIS_DB=0
REDIS_PASSWORD=
```

Change them if you want. Default rate limit is 5 requests per 5 seconds (configured in code).

## Running

```bash
go run cmd/api/main.go
```

Or build it:
```bash
go build -o bin/api cmd/api/main.go
./bin/api
```

## Testing it out

Just hit the API a bunch of times:

```bash
curl http://localhost:8081/health
```

Do it more than 5 times in 5 seconds and you'll get rate limited.

## Project structure

```
├── cmd/
│   └── api/           # main app and HTTP stuff
├── internal/
│   ├── env/           # env variable helpers
│   └── ratelimiter/   # rate limiter implementation
├── docker-compose.yml # Redis setup
└── .env              # configuration
```

## Redis Commander

There's a Redis Commander instance in the docker-compose if you want to peek at what's in Redis. It's at `http://localhost:8081` after you start the containers.

## TODO

- [ ] Add more rate limiting algorithms (sliding window, token bucket)
- [ ] Better error responses when rate limited
- [ ] Add some actual tests
- [ ] Maybe add prometheus metrics?

## Notes

This was mainly a learning project to understand rate limiting better. It works fine but probably needs more polish before using in production. Feel free to fork and improve it.

## License

MIT - do whatever you want with it.
