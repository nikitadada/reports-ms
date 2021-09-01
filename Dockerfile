FROM alpine:3.10

RUN apk update
COPY ./report-action /usr/bin/report-action

CMD ["report-action"]
