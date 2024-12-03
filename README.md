# ogp-fetch
## Dev
URL in `$1` Args
```bash
go build -o ogp-fetch
./ogp-fetch https://soulminingrig.com/
```
or
```bash
cat urls | while read -r url ; do ./ogp-fetch $url ; done 
```
```bash
cat urls | xargs -P 10 -I {} ./ogp-fetch {}
```
