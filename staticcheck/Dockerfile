FROM golang:1.19 AS builder
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

FROM golang:1.19 AS runtime
COPY --from=builder /go/bin/staticcheck /go/bin/staticcheck
