FROM golang:1.20.3-alpine as builder
ARG SERVICE_NAME
ARG BRANCH_NAME
ARG BUILD_NAME
ARG SEALIGHTS_TOKEN
ENV OS_ARCH linux-amd64
ADD . go/src
WORKDIR go/src
RUN go mod download
### Sealights instrumentation
RUN wget -U "slgoagent" -q -O slgoagent.tar.gz \
   https://agents.sealights.co/slgoagent/latest/slgoagent-${OS_ARCH}.tar.gz \
    && tar -xvf slgoagent.tar.gz \
    && rm slgoagent.tar.gz
RUN wget -U "slcli" -q -O slcli.tar.gz \
    https://agents.sealights.co/slcli/latest/slcli-${OS_ARCH}.tar.gz \
    && tar -xvf slcli.tar.gz \
    && rm slcli.tar.gz
RUN echo ${SEALIGHTS_TOKEN} >> ./token.txt
RUN ./slcli config init --lang go --token ./token.txt
RUN ./slcli config create-bsid --app ${SERVICE_NAME} --branch ${BRANCH_NAME} --build ${BUILD_NAME}
RUN ./slcli scan --bsid buildSessionId.txt --path-to-scanner ./slgoagent --workspacepath ./ --scm git --scmProvider github
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -mod=vendor -o ./go-calc-demo ./service/cmd/...
CMD ["./go-calc-demo"]
