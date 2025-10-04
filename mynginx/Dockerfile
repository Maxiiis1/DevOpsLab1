FROM alpine:3.18

RUN apk add --no-cache nginx

RUN adduser -D -H -u 1001 user

RUN mkdir -p /run/nginx /usr/share/nginx/html /tmp/nginx \
    && chown -R user:user /run/nginx /usr/share/nginx /etc/nginx /tmp/nginx

COPY nginx.conf /etc/nginx/nginx.conf

USER user

EXPOSE 8080

CMD ["nginx", "-g", "error_log /dev/stderr info; daemon off;"]
