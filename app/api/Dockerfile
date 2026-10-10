# =============== #
# Stage 1: Build  #
# =============== #
from golang:1.26.2 as builder

workdir /builder

copy go.mod go.sum .

run go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd


# =================== #
# Stage 2: Production #
# =================== #
from golang:1.26.2

workdir /app

COPY --from=build /server ./server

EXPOSE 9000

CMD ["./server"]
