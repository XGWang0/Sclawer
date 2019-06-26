# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /app
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV GOPATH=/go/src/github.com/XGWang0/Sclawer
ENV SRC_DIR=$GOPATH/src
ENV TEMPLATE_FILE=$SRC_DIR/template
ENV STATIC_FILE=$SRC_DIR/static

# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN cd $SRC_DIR; go build -o myapp; cp myapp /app/;cp -r $TEMPLATE_FILE /app/; cp -r $STATIC_FILE /app/
ENTRYPOINT ["./myapp"]
