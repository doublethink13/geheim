FROM golang:1.17.5

# get -v
RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
    go install github.com/ramya-rao-a/go-outline@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest

COPY .ssh /root/.ssh

RUN chmod 400 /root/.ssh/id* /root/.ssh/config && \
    chown -R root:root /root/.ssh

COPY .bashrc /root
