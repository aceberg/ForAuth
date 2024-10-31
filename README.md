[![Main-Docker](https://github.com/aceberg/forauth/actions/workflows/main-docker.yml/badge.svg)](https://github.com/aceberg/forauth/actions/workflows/main-docker.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/forauth)](https://goreportcard.com/report/github.com/aceberg/forauth)
[![Maintainability](https://api.codeclimate.com/v1/badges/e8f67994120fc7936aeb/maintainability)](https://codeclimate.com/github/aceberg/forauth/maintainability)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/aceberg/forauth)

<h1><a href="https://github.com/aceberg/forauth">
    <img src="https://raw.githubusercontent.com/aceberg/forauth/main/assets/logo.png" width="35" />
</a>ForAuth</h1>

ForAuth (Forward Auth) - simple reverse proxy with session-cookie auth and notifications on login.

![Screenshot](https://raw.githubusercontent.com/aceberg/forauth/main/assets/Screenshot.png)


## Config


Configuration can be done through config file or environment variables

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| FA_AUTH | Enable Session-Cookie authentication | false |
| FA_AUTH_EXPIRE | Session expiration time. A number and suffix: **m, h, d** or **M**. | 7d |
| FA_AUTH_USER | Username |  |
| FA_AUTH_PASSWORD | Encrypted password (bcrypt). [How to encrypt password with bcrypt?](docs/BCRYPT.md) |  |

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| FA_HOST | Listen address for both Config and Proxy | 0.0.0.0 |
| FA_PORT   | Port for Proxy | 8800 |
| FA_PORTCONF   | Port for Config page | 8801 |
| FA_TARGET   | Where to proxy after login (host:port). Example: `0.0.0.0:8840` |  |
| FA_THEME | Any theme name from https://bootswatch.com in lowcase or [additional](https://github.com/aceberg/aceberg-bootswatch-fork) (emerald, grass, grayscale, ocean, sand, wood)| united |
| FA_COLOR | Background color: light or dark | dark |
| FA_NODEPATH   | Path to local JS and Themes ([node-bootstrap](https://github.com/aceberg/my-dockerfiles/tree/main/node-bootstrap)) |  |
| FA_NOTIFY   | Shoutrrr URL. ForAuth uses [Shoutrrr](https://github.com/containrrr/shoutrrr) to send notifications. It is already integrated, just needs a correct URL. Examples for Discord, Email, Gotify, Matrix, Ntfy, Pushover, Slack, Telegram, Generic Webhook and etc are [here](https://containrrr.dev/shoutrrr/v0.8/services/gotify/) |  |
| TZ | Set your timezone for correct time |  |

## Options

| Key  | Description | Default | 
| --------  | ----------- | ------- | 
| -d | Path to config dir | /data/ForAuth | 
| -n | Path to local JS and Themes ([node-bootstrap](https://github.com/aceberg/my-dockerfiles/tree/main/node-bootstrap)) |  | 

## Local network only
By default, this app pulls themes, icons and fonts from the internet. But, in some cases, it may be useful to have an independent from global network setup. I created a separate [image](https://github.com/aceberg/my-dockerfiles/tree/main/node-bootstrap) with all necessary modules and fonts.    
```sh
docker run --name node-bootstrap       \
    -v ~/.dockerdata/icons:/app/icons  \ # For local images
    -p 8850:8850                       \
    aceberg/node-bootstrap
```
```sh
docker run --name exdiary \
    -v ~/.dockerdata/forauth:/data/forauth \
    -p 8851:8851 \
    aceberg/forauth -n "http://$YOUR_IP:8850"
```
Or use [docker-compose](docker-compose-local.yml)



## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/forauth/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)
- Favicon and logo: [Flaticon](https://www.flaticon.com/icons/)