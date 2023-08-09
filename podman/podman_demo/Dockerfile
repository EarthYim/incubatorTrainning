FROM alpine 
RUN apk update && apk add go
COPY . .
CMD go build main.go
# ENTRYPOINT /main