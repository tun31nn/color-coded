# Go 
FROM alpine:latest
RUN apk update; apk upgrade
# Copy binary to app folder
RUN adduser -D -h /home/color-coded color-coded
WORKDIR /home/color-coded
COPY bin/color-coded .
RUN chown color-coded:color-coded -R /home/color-coded
USER color-coded
EXPOSE 8080
ENTRYPOINT ["/home/color-coded/color-coded"]
