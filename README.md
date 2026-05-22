# otp_breaker

`otp_breaker` is a small Go CLI that generates numeric OTP candidates and sends them to an HTTP endpoint using multiple workers.

Use it only against systems you own or are explicitly authorized to test.

## What it does

- Generates zero-padded OTP values of a fixed length.
- Sends each candidate to the configured URL.
- Runs requests in parallel with a worker pool.
- Prints the HTTP status code and response body for each request.

The main entrypoint is in [cmd/internal/main.go](cmd/internal/main.go).

## Requirements

- Go 1.25.4 or newer

## Usage

Run the tool with `go run`:

```bash
go run ./cmd/internal -url https://example.com/reset -method POST -t 10 -l 4
```

You can also build a binary first:

```bash
go build -o otp_breaker ./cmd/internal
./otp_breaker -url https://example.com/reset -method POST -t 10 -l 4
```

## Flags

- `-url`: Target endpoint URL.
- `-method`: HTTP method to use. Default: `POST`.
- `-t`: Number of worker goroutines. Default: `10`.
- `-l`: OTP length. Default: `4`.

## Example

```bash
go run ./cmd/internal \
	-url https://example.com/api/reset-password \
	-method POST \
	-t 20 \
	-l 6
```

## Notes

- Requests are sent with a JSON body containing the generated `token` value.
- The current implementation also sets a fixed `Host` header and includes placeholder password fields in the request body.
