FROM registry.hundsun.com/library/busybox:1.29.2

COPY web_echo /
RUN chmod +x /web_echo
EXPOSE 8080

CMD ["/web_echo"]
