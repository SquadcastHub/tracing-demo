FROM golang:1.20

RUN mkdir /app
COPY . /app/

RUN cd /app && go build -o backend.bin -ldflags="-X 'main.version=${VERSION}'"
EXPOSE 8000
ENTRYPOINT ["/app/backend.bin"]
