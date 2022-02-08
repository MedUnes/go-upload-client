# go-upload-client

Client is used to upload files to [Myra's](https://upload.myracloud.com/en/) PushCDN.
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

* The client requires a configuration file to operate.
Please see the example:

```yaml
endpoint: https://upload.myracloud.com
language: en
proxy: "socks5://localhost:8080"

login:
    user: user1
    apikey: deadbeef1fbd41116fe98bbeeeeeeeef
    secret: deadaef1fbd41116fe98beefabc14142
```

* If you do not use a proxy leave the field empty.
To get the login data you need to contact [Myra Support](https://myrasecurity.com/en/).

* The client will look up the config in the following order:

* Config given via commandline flag
* $HOME/config.yml
* ./config.yml

* You can also copy the sample config file:
```bash
user@host:~/go-upload-client$ cp config.yml.dist config.yml
```
then replace placeholders (like: *__MYRA_CDN_USER__* ) inside the ```config.yml``` 
with your real values.

## Usage

See help for usage instructions.

    ./myra-upload --help

## Docker
* It is also possible to have the Go Upload Client up and running on a dockerized environment. 
This can save you time and effort you'd have spent setting up Golang and building the application.
### Building the image
* First off, you have to build the docker image. It is as easy as running the following command:
```bash
user@host:~/go-upload-client$ make docker-build
```

### Using the image
* Once you have built the image,  you are ready to go. Login to the container:
```bash
user@host:~/go-upload-client$ make docker-login
```
and from there, you can run the myra-upload command:
```bash
root@463c419570cd:/go/src/github.com/Myra-Security-GmbH/go-upload-client# ./myra-upload --help
Usage: myra-upload [--init] --domain DOMAIN --bucket BUCKET [--recursive] [--silent] [--configfile CONFIGFILE] SOURCE TARGET

Positional arguments:
  SOURCE                 Source file or folder
  TARGET                 Target folder

Options:
  --init                 Creates a default configuration file
  --domain DOMAIN, -d DOMAIN
                         Domain
  --bucket BUCKET, -b BUCKET
                         Bucket
  --recursive, -r        Upload folder recursive
  --silent, -s           No progress output
  --configfile CONFIGFILE, -c CONFIGFILE
                         Configfile to use [default: ./config.yml]
  --help, -h             display this help and exit
```

Then of course you can do the upload from within the container.
### Stopping and spinning up the container

* If you want to stop the container (ex: save resources), just run the command:
```bash
user@host:~/go-upload-client$ make docker-stop
```

* If you want to spin it up again (without rebuilding it again):
```bash
user@host:~/go-upload-client$ make docker-up
```
