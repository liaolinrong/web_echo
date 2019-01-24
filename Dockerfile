FROM registry.hundsun.com/library/busybox:1.28.0

COPY web_echo /
RUN chmod +x /web_echo

CMD ["/web_echo"]
