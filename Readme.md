# Stellar Exercise

What is this? 

Yet another interview coding challenge solution that I coded up in a jiffy in off hours. Below are the requirements given. Yes the challenge was timed with a very short deadline. A $100 Amazon gift card was offered. I think it pretty much achieved "minimum viable product" for the requirements given and as quickly as possible. I don't think it's too bad myself. The organization did make good on the gift card.

### This was the requirement given:
 

# Stellar Coding Challenge

## Summary

Write an API web server that accepts a snippet of text and makes that snippet
available at a URL. Each snippet should be available as text at a URL until it
expires. A snippet's expiry should be extended by 30 seconds each time it is accessed.

The request to store the snippet should accept this information:

| Name       | Description                                      |
|------------|--------------------------------------------------|
| name       | Name of the snippet                              |
| expires_in | Seconds from now until the snippet should expire |
| snippet    | Contents of the snippet                          |

The request to store the snippet should be replied to with a response that
includes the URL where the snippet can be read.

Snippets can be stored in memory, and do not need to be editable after storing.
The solution needs only to be an API, not a graphical or website user
interface.

Snippets should be retrievable by name.

## What we're looking for

* An actual HTTP web server that runs and can be accessed through a URL
* A clean, minimalistic implementation. Focus on the core functionality, and pay extra attention to API the service should serve
* Appropriate use of web frameworks and open-source libraries as necessary. Think of this as building an MVP
* Too little time? Your prioritization skills are also being evaluated
* Readable and clear code, so that everyone and anyone can understand it

## What we're not looking for

* Don't write your own HTTP implementation. You won't get extra points for it, and this will probably take time that could be spent elsewhere
* This is a simple problem. If your solution is complex, take a step back
* Do not reinvent the wheel. Make use of all the tools you know and like, as well as the knowledge you've built throughout your career

## Test Cases

We expect you to implement this API exactly. Please ensure your solution handles these test cases correctly.

```sh

curl -X POST -H "Content-Type: application/json" -d '{"name":"recipe", "expires_in": 30, "snippet":"1 apple"}' https://example.com/snippets
# response 201 Created
{
  "url": "https://example.com/snippets/recipe",
  "name": "recipe",
  "expires_at": "2020-02-22T20:02:02Z",
  "snippet": "1 apple"
}

curl https://example.com/snippets/recipe
# response 200 OK
{
  "url": "https://example.com/snippets/recipe",
  "name": "recipe",
  "expires_at": "2020-02-22T20:02:32Z",
  "snippet": "1 apple"
}

# wait 60 seconds

curl https://example.com/snippets/recipe
# response 404 Not Found

```

## Instructions I provided for the evaluator of this solution:

### 1. Install gorilla mux package
```bash
go get -u github.com/gorilla/mux
```
### 3. Build the project
```bash
go build
```
### 3. Run the web server passing flags for cert and key files (you will need your own cert and key files in your own environment)
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
