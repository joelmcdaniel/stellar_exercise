# Stellar Exercise

### 1. Install the following packages
```bash
go get -u github.com/gorilla/mux
```
```bash
go get github.com/subosito/gotenv
```
### 2. Set the HOST and PORT environment variables in .env
```bash
HOST = localhost
PORT = 8080
```
### 3. Build the project
```bash
go build
```
### 3. Start the web server
```bash
./stellar_exercise
```
### 4. Run test cases with curl
*Please note*: I ran into an issues with curl running https on my development machine. I attempted to resolve,
however due to lack of time the project will expect all curl request to be http only at http://[host]:[port].
#### Create snippet
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"recipe", "expires_in": 30, "snippet":"1 apple"}' http://localhost:8080/snippets
# response 201 Created
{
  "url": "http://localhost:8080/snippets/recipe",
  "name": "recipe",
  "expires_at": "2020-12-13T20:02:02Z",
  "snippet": "1 apple"
}
```
#### Fetch snippet
```bash
curl http://localhost:8080/snippets/recipe
# response 200 OK
{
  "url": "http://localhost:8080/snippets/recipe",
  "name": "recipe",
  "expires_at": "2020-12-13T20:02:32Z",
  "snippet": "1 apple"
}
```
#### Wait 60 seconds and try to fetch again
```bash
curl http://localhost:8080/snippets/recipe
# response 404 Not Found
```