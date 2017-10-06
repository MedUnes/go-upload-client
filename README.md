# go-upload-client

Client is used to upload files to [Myra's](https://myracloud.com/en/) PushCDN.
In order to use this client you need to be a customer of Myra.

## Build

### Requirements

* https://github.com/Masterminds/glide
* make
* https://upx.github.io (optional)

### Generate 

    glide install
    make

## Configuration

The client requires a configuration file to operate.
Please see the example:

```yaml
endpoint: http://upload.myracloud.com
language: en
proxy: "socks5://localhost:8080"

login:
    user: user1
    apikey: deadbeef1fbd41116fe98bbeeeeeeeef
    secret: deadaef1fbd41116fe98beefabc14142
```

If you do not use a proxy leave the field empty.
To get the login data you need to contact [Myra Support](https://myracloud.com/en/).

The client will lookup the config in the following order:

* Config given via commandline flag
* $HOME/config.yml
* ./config.yml

## Usage

See help for usage instructions.

    ./myra-upload --help
