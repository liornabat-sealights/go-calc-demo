FROM golang:1.20.3-alpine as builder
ARG SERVICE_NAME
ARG BRANCH_NAME
ARG BUILD_NAME
ARG SEALIGHTS_TOKEN
ENV OS_ARCH linux-amd64

# Echo build arguments for debugging
RUN echo "SERVICE_NAME: ${SERVICE_NAME}"
RUN echo "BRANCH_NAME: ${BRANCH_NAME}"
RUN echo "BUILD_NAME: ${BUILD_NAME}"
RUN echo "SEALIGHTS_TOKEN: ${SEALIGHTS_TOKEN}"

# Add the source code
ADD . go/src
WORKDIR go/src

# Download dependencies
RUN go mod download

### Sealights instrumentation
# Download and extract Sealights Go agent
RUN wget -U "slgoagent" -q -O slgoagent.tar.gz \
   https://agents.sealights.co/slgoagent/latest/slgoagent-${OS_ARCH}.tar.gz \
    && tar -xvf slgoagent.tar.gz \
    && rm slgoagent.tar.gz

# Download and extract Sealights CLI
RUN wget -U "slcli" -q -O slcli.tar.gz \
    https://agents.sealights.co/slcli/latest/slcli-${OS_ARCH}.tar.gz \
    && tar -xvf slcli.tar.gz \
    && rm slcli.tar.gz

# Save the Sealights token to a file and print its contents
RUN echo "${SEALIGHTS_TOKEN}" > ./token.txt \
    && cat ./token.txt

# Sealights configuration
RUN ./slcli config init --lang go --token ./token.txt
RUN ./slcli config create-bsid --app ${SERVICE_NAME} --branch ${BRANCH_NAME} --build ${BUILD_NAME}

# Run Sealights scan
RUN ./slcli scan --bsid buildSessionId.txt --path-to-scanner ./slgoagent --workspacepath ./ --scm git --scmProvider github

# Run tests
RUN go test -v ./...

# Build the Go application
RUN go mod vendor

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -mod=vendor -o ./go-calc-demo

# Define the entry point for the container
CMD ["./go-calc-demo"]