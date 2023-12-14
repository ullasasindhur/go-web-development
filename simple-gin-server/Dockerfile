FROM golang:1.21 as build-stage
WORKDIR /app 
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /simple-gin-server


FROM gcr.io/distroless/base-debian11 as release-stage
WORKDIR /
COPY --from=build-stage /simple-gin-server /simple-gin-server
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT [ "/simple-gin-server" ]