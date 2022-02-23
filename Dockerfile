FROM golang:1.15

# We need to add ${GOPATH}/bin to PATH to have access dlv
ENV PATH ${GOPATH}/bin:$PATH
ENV GO111MODULE=on

# Setting working directory
WORKDIR ${GOPATH}/src/github.com/RexterR/imger

# Copy source code
COPY . .

# Get delve debugger and gin code reloader
RUN go get github.com/go-delve/delve/cmd/dlv
RUN go get github.com/codegangsta/gin

# Build our application
RUN go build -o imgerapi cmd/main.go

# Expose server and debug port
EXPOSE 4005 2345

# Execute our application
CMD ["./imgerapi"]