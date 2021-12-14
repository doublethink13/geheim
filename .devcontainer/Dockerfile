FROM golang:1.17.5

RUN go get -v golang.org/x/tools/gopls@latest && \
    go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
    go get -v github.com/ramya-rao-a/go-outline@latest && \
    go get -v github.com/go-delve/delve/cmd/dlv@master && \
    go get -v github.com/go-delve/delve/cmd/dlv@latest && \
    go get -v honnef.co/go/tools/cmd/staticcheck@latest


COPY .ssh /root/.ssh
RUN chmod 400 /root/.ssh/id* /root/.ssh/config && \
    chown -R root:root /root/.ssh

COPY .bashrc /root
