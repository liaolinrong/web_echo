# web_echo
web server for echo host IP

docker run command:
docker run -it -e LISTEN_PORT=8123 -p 8123:8123 --rm registry.hundsun.com/hcs/web_echo:0.1

k8s yaml:
web_echo.yaml
