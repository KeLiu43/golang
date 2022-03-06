FROMgolang:1.17AS build
WORKDIR /httpserver/
COPY . .
ENVCGO_ENABLED=0
ENVGO111MODULE=on
ENVGOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

FROMbusybox
COPY --from=build /httpserver/httpserver /httpserver/httpserver
EXPOSE8360
ENVENVlocal
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]