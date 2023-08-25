FROM golang:1.20

RUN mkdir /app
COPY . /app/

RUN cd /app && go build -o frontend.bin -ldflags="-X 'main.version=${VERSION}'"
EXPOSE 8080
ENTRYPOINT ["/app/frontend.bin"]
