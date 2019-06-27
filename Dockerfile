# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /app
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name

# Add the source code:
ADD sclawer  /app/
ADD template/ /app/template/
ADD static/ /app/static/
# Build it:
ENTRYPOINT ["./sclawer"]
