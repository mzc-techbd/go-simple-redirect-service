```powershell
$env:REDIRECT_SERVICE_URL="https://google.com"
```

```powershell

docker build -t asia-northeast3-docker.pkg.dev/khc-common/kktae-tmp/go-redirect-service .

docker push asia-northeast3-docker.pkg.dev/khc-common/kktae-tmp/go-redirect-service

```