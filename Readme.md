# Stellar Exercise

### 1. Install gorilla mux package
```bash
go get -u github.com/gorilla/mux
```
### 3. Build the project
```bash
go build
```
### 3. Run the web server passing flags for cert and key files
```bash
[sudo] ./stellar_exercise -cert=<path to cert file> -key=<path to key file>
```
### 4. Run test cases with curl
#### Create snippet
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"recipe", "expires_in": 30, "snippet":"1 apple"}' https://example.com/snippets
# response 201 Created
{
  "url": "https://example.com/snippets/recipe",
  "name": "recipe",
  "expires_at": "2020-12-13T20:02:02Z",
  "snippet": "1 apple"
}
```
#### Fetch snippet
```bash
curl https://example.com/snippets/recipe
# response 200 OK
{
  "url": "https://example.com/snippets/recipe",
  "name": "recipe",
  "expires_at": "2020-12-13T20:02:32Z",
  "snippet": "1 apple"
}
```
#### Wait 60 seconds and try to fetch again
```bash
curl https://example.com/snippets/recipe
# response 404 Not Found
```