[![Main-Docker](https://github.com/aceberg/ForAuth/actions/workflows/main-docker-all.yml/badge.svg)](https://github.com/aceberg/ForAuth/actions/workflows/main-docker-all.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/forauth)](https://goreportcard.com/report/github.com/aceberg/forauth)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/aceberg/forauth)

<h1><a href="https://github.com/aceberg/forauth">
    <img src="https://raw.githubusercontent.com/aceberg/forauth/main/assets/logo.png" width="35" />
</a>ForAuth</h1>

ForAuth (Forward Auth) - simple auth app (session-cookie) with notifications on login and [multiple targets and users](https://github.com/aceberg/forauth#multiple-targets-and-users) option

- [Security](https://github.com/aceberg/forauth#security)
- [Quick start](https://github.com/aceberg/forauth#quick-start)
- [Config](https://github.com/aceberg/forauth#config)
- [Options](https://github.com/aceberg/forauth#options)
- [Multiple Targets and Users](https://github.com/aceberg/forauth#multiple-targets-and-users)
- [Local network only](https://github.com/aceberg/forauth#local-network-only)
- [CURL](https://github.com/aceberg/forauth#curl)
- [Thanks](https://github.com/aceberg/forauth#thanks)

![Screenshot](https://raw.githubusercontent.com/aceberg/forauth/main/assets/Screenshot.png)    
<details>
  <summary>Screenshot 2</summary>

![Screenshot1](https://raw.githubusercontent.com/aceberg/forauth/main/assets/Screenshot1.png)
</details>

## Securuty
- This app is only safe when used with `https`
- Use strong password
- Make sure direct access to Target app is closed with firewall or other measures

## Quick start
```sh
docker run --name forauth \
    -v ~/.dockerdata/ForAuth:/data/ForAuth \
    -p 8800:8800 \ # Proxy port
    -p 8801:8801 \ # Config port
    aceberg/forauth
```
Then open Config page in browser and set up Auth and Target app.   

Example [docker-compose-auth.yml](https://github.com/aceberg/WatchYourPorts/blob/main/docker-compose-auth.yml) for [WatchYourPorts](https://github.com/aceberg/WatchYourPorts). This should work with other apps too.

## Config

Configuration can be done through config file, GUI or environment variables. Variable names is `config.yaml` file are the same, but in lowcase.

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| FA_AUTH | Enable Session-Cookie authentication | false |
| FA_AUTH_EXPIRE | Session expiration time. A number and suffix: **m, h, d** or **M**. | 7d |
| FA_AUTH_USER | Main user username |  |
| FA_AUTH_PASSWORD | Encrypted password (bcrypt). [How to encrypt password with bcrypt?](docs/BCRYPT.md) |  |

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| FA_HOST | Listen address for both Config and Proxy | 0.0.0.0 |
| FA_PORT   | Port for Proxy | 8800 |
| FA_PORTCONF   | Port for Config page | 8801 |
| FA_TARGET   | Where to proxy after login (host:port). Example: `192.168.1.1:8840` |  |
| FA_THEME | Any theme name from https://bootswatch.com in lowcase or [additional](https://github.com/aceberg/aceberg-bootswatch-fork) (emerald, grass, grayscale, ocean, sand, wood)| united |
| FA_COLOR | Background color: light or dark | dark |
| FA_NODEPATH   | Path to local JS and Themes ([node-bootstrap](https://github.com/aceberg/my-dockerfiles/tree/main/node-bootstrap)) |  |
| FA_NOTIFY   | Shoutrrr URL. ForAuth uses [Shoutrrr](https://github.com/containrrr/shoutrrr) to send notifications. It is already integrated, just needs a correct URL. Examples for Discord, Email, Gotify, Matrix, Ntfy, Pushover, Slack, Telegram, Generic Webhook and etc are [here](https://containrrr.dev/shoutrrr/v0.8/services/gotify/) |  |
| FA_NOTIFY2 | Second Shoutrrr URL. The app will send notifications to both, if they are not empty | |
| FA_IPINFO | Get client IP info (from https://ipinfo.io) on login | false |
| TZ | Set your timezone for correct time |  |

## Options

| Key  | Description | Default | 
| --------  | ----------- | ------- | 
| -d | Path to config dir | /data/ForAuth | 
| -n | Path to local JS and Themes ([node-bootstrap](https://github.com/aceberg/my-dockerfiles/tree/main/node-bootstrap)) |  | 

## Multiple Targets and Users
Multiple Targets and Users for each target can be configured from `Advanced` page or in `targets.yaml` file inside the config dir. Main user (`FA_AUTH_USER`) has access to all targets and config.

<details>
  <summary>Example:</summary>

```yaml
0.0.0.0:8854:               # where proxy will listen
    name: DiaryMD           # name
    target: 127.0.0.1:8754  # where an app listens
    users:                  # users of this target
        user1:              # username
            enabled: true   # must be true for user to login
            username: user1 # username (same as above)
            password: $2a$10$bPH6208LpuJFos3x1VhFA.PxzygaAhT056uPxspJxwccgP4n.AnEe
            expire: 14d     # session expiration time
        user3:
            enabled: true
            username: user3
            password: $2a$10$eZp3I0A9ojT32gTXvPscHec9e7cHHYtb6M6phl2mUdHXyhFosLW.C
            expire: 1d
0.0.0.0:8855:
    name: AnyAppStart
    target: 127.0.0.1:8755
    # users:                  # users section is optional
                              # without it only Main user can login
```

</details>

## Local network only
By default, this app pulls themes, icons and fonts from the internet. But, in some cases, it may be useful to have an independent from global network setup. I created a separate [image](https://github.com/aceberg/my-dockerfiles/tree/main/node-bootstrap) with all necessary modules and fonts.    
```sh
docker run --name node-bootstrap       \
    -p 8850:8850                       \
    aceberg/node-bootstrap
```
```sh
docker run --name forauth \
    -v ~/.dockerdata/ForAuth:/data/ForAuth \
    -p 8800:8800 \
    -p 8801:8801 \
    aceberg/forauth -n "http://$YOUR_IP:8850"
```
## CURL
To access Target app with `curl`:

```sh
curl -X POST http://localhost:8800 -H "Content-Type: application/x-www-form-urlencoded" -d "username=user&password=pw" -c fileCookie
```
```sh
curl http://localhost:8800 -b fileCookie
```

## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/forauth/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)
- Favicon and logo: [Flaticon](https://www.flaticon.com/icons/)