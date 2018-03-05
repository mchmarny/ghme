FROM golang:1.9.2
WORKDIR /go/src/github.com/mchmarny/github-teams-utils/
COPY . .

# restore to pinnned versions of dependancies 
RUN go get github.com/tools/godep
RUN godep restore

RUN CGO_ENABLED=0 GOOS=linux go build .


FROM scratch
COPY --from=0 /go/src/github.com/mchmarny/github-teams-utils/github-teams-utils .
ENTRYPOINT ["/github-teams-utils"]
