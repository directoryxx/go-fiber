FROM golang:buster

# Install VSCode Dependencies Suggestion
RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs
RUN go get github.com/ramya-rao-a/go-outline
RUN go get github.com/cweill/gotests/gotests
RUN go get github.com/fatih/gomodifytags
RUN go get github.com/josharian/impl
RUN go get github.com/haya14busa/goplay/cmd/goplay
RUN go get github.com/go-delve/delve/cmd/dlv
RUN go get github.com/go-delve/delve/cmd/dlv@master
RUN go get honnef.co/go/tools/cmd/staticcheck
RUN go get golang.org/x/tools/gopls

RUN mkdir /deploy
WORKDIR /deploy

COPY . /deploy

RUN go mod download
