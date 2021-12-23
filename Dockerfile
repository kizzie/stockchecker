FROM golang:1.17.5-bullseye as build
COPY . /app
WORKDIR /app
RUN go get . && go build -o app . 

FROM golang:1.17.5-bullseye as final

ENV GIN_MODE=release
ENV PORT=8080

RUN mkdir /app && useradd app 
COPY --from=build /app/app /app/stockchecker
RUN chown -R app:app /app
USER app
EXPOSE 8080

CMD ["/app/stockchecker"]

