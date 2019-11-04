

all: mac linux64 linux32

mac: 
	GOOS=darwin GOARCH=amd64 go build  -o binaries/mac/gohttpserver

linux64:
	GOOS=linux GOARCH=amd64 go build  -o binaries/linux64/gohttpserver assets_vfsdata.go  httpstaticserver.go  ipa.go  main.go  oauth2-proxy.go  openid-login.go  res.go  utils.go xjconvert.go  zip.go  


linux32:
	GOOS=linux GOARCH=386 go build  -o binaries/linux386/gohttpserver
