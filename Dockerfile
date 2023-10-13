FROM golang:alpine AS build
WORKDIR /
COPY . .
RUN go build -o /bin/app main.go

FROM scratch
EXPOSE 8080
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
