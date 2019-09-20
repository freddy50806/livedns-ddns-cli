# livedns-ddns-cli
DDNS client writes with liveDNS api to update your IP address to DNS 

## Usage
require go.12 above
### Build
```
git clone git@github.com:freddy50806/livedns-ddns-cli.git
cd livedns-ddns-cli
go mod tidy
go build
```
### Add your infomation
Fill you `API Key` and `Domain Name` in `config.json`, you can set request interval by `duration` field(default is 30 seconds),which not be shorter than 30 seconds.
```
##config.json
{
    "api_key":"<YOU_API_KEY>",
    "domain":"<YOUR_DOMAIN_NAME>",
    "duration":"30s"
}
```

## DockerFile
### Build Image
Fill your info in `config.json` and run command below build the image and run the container `ddns-cli` on your server.
```
docker build -t=livedns .
docker run -d --name ddns-cli livedns
```
### Modify Configure 
Write new `config.json` and replace to docekr `config.json`
```
docker cp ./config.json ddns-cli:/opt/livedns-ddns-cli/
```