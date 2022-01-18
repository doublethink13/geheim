FROM golang:1.17.5

RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0 && \
    golangci-lint --version

COPY .ssh /root/.ssh

RUN chmod 400 /root/.ssh/id* /root/.ssh/config && \
    chown -R root:root /root/.ssh

COPY .bashrc /root
