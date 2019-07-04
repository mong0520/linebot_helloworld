## Steps

Install docker

Install docker-copmse
```shell
sudo curl -L "https://github.com/docker/compose/releases/download/1.24.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```

Build
```
make build
```

Start
```
docker-compose up
```

Register URL `https://b64ae032.ngrok.io/callback` to LINE message API webhook
```
ngrok_1    | The ngrok tunnel is active
ngrok_1    | https://b64ae032.ngrok.io ---> linebot:2000
```

Rebuild and restart
```
make build
docker-compose down linebot
docker-compose up linebot

docker ps
```
CONTAINER ID        IMAGE                     COMMAND             CREATED             STATUS              PORTS                    NAMES
b78ef77f6fcd        gtriggiano/ngrok-tunnel   "npm start"         2 minutes ago       Up 8 seconds        4040/tcp                 linebot_helloworld_ngrok_1
66e9cdf48983        linebot_helloworld        "/bin/sh -c /app"   2 minutes ago       Up 8 seconds        0.0.0.0:2000->2000/tcp   linebot_helloworld_linebot_1
```