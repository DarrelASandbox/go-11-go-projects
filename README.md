<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#01-simple-http-server">01-simple-http-server</a></li>
    <li><a href="#02-movies-crud">02-movies-crud</a></li>
    <li><a href="#03-book-management-api">03-book-management-api</a></li>
    <li><a href="#04-slack-bot-calculates-age">04-slack-bot-calculates-age</a></li>
    <li><a href="#05-slack-bot-uploads-file">05-slack-bot-uploads-file</a></li>
    <li><a href="#06-email-verifier-tool">06-email-verifier-tool</a></li>
    <li><a href="#07-simple-aws-lambda-example">07-simple-aws-lambda-example</a></li>
  </ol>
</details>

&nbsp;

## About The Project

- [Learn Go Programming by Building 11 Projects – Full Course](https://www.youtube.com/watch?v=jFfo23yIWac)
- [freeCodeCamp](https://www.freecodecamp.org/)
- [Akhil Sharma](https://github.com/AkhilSharma90)

1. [Original Repo: simple-http-server](https://github.com/AkhilSharma90/simple-http-server-GO)
2. [Original Repo: movies-crud](https://github.com/AkhilSharma90?tab=repositories&type=source)
3. [Original Repo: book-management-api](https://github.com/AkhilSharma90/Golang-MySQL-CRUD-Bookstore-Management-API)
4. [Original Repo: slack-bot-calculates-age](https://github.com/AkhilSharma90/GO-Slackbot-Calculates-Age)
5. [Original Repo: slack-bot-uploads-file](https://github.com/AkhilSharma90/GO-SlackBot-Uploads-File)
6. [Original Repo: email-verifier-tool](https://github.com/AkhilSharma90/GO-Email-Checking-Tool)
7. [Original Repo: simple-aws-lambda-example](https://github.com/AkhilSharma90/Simple-go-AWS-Lambda-example)
8. [Original Repo: beginner-CRM-project](https://github.com/AkhilSharma90/go-beginner-CRM-project)
9. [Original Repo: beginner-fiber-HRMS-project](https://github.com/AkhilSharma90/go-beginner-fiber-HRMS-project)
10. [Original Repo: serverless-project](https://github.com/AkhilSharma90/Golang-Serverless-Project)
11. [Original Repo: ai-bot-wolfram-slack](https://github.com/AkhilSharma90/AI-Bot-GOlang-Wit.ai-Wolfram-Slack)

&nbsp;

---

&nbsp;

## 01-simple-http-server

```
        -> /      -> index.html
Server  -> /hello -> hello func
        -> /form  -> form func  -> form.html
```

&nbsp;

---

&nbsp;

## 02-movies-crud

- No database
- CRUD
- Gorilla Mux

```
                                          ROUTES        FUNCTIONS     ENDPOINTS       METHODS

                                        -> GET ALL      getMovies     -> /movies      -> GET      <--|
                                        -> GET BY ID    getMovie      -> /movies/id   -> GET      <--|
Movies Server serving using Gorilla MUX -> CREATE       creteMovie    -> /movies      -> POST     <--- POSTMAN
                                        -> UPDATE       updateMovie   -> /movies/id   -> PUT      <--|
                                        -> DELETE       deleteMovie   -> /movies/id   -> DELETE   <--|
```

&nbsp;

---

&nbsp;

## 03-book-management-api

- MySQL
- GORM
- json Marshal & Unmarshal
- Project structure
- Gorilla Mux

```
CMD -> main.go

    -> config       -> app.go
    -> controllers  -> book controllers
PKG -> models       -> book.go
    -> routes       -> bookstore routes
    -> utils        -> utils.go


CONTROLLERS
GET     -> /book            -> GETBOOKS
POST    -> /book            -> CREATE BOOK
GET     -> /book/{bookId}   -> GET BOOK BY ID
PUT     -> /book/{bookId}   -> UPDATE BOOK
DELETE  -> /book/{bookId}   -> DELETE BOOK
```

&nbsp;

---

&nbsp;

## 04-slack-bot-calculates-age

- `.env`:
  - `SLACK_BOT_TOKEN=`
  - `SLACK_SOCKET_TOKEN=`
- [slack api](https://api.slack.com/)
- Create slack workspace -> Create app
- Connect using Socket Mode -> Generate socket-token
- Event Subscription -> Enable Events -> Subscribe to bot events -> Add Bot User Event
  - app_mention
  - im_history_changed
  - message.im
  - message.channels
  - message.mpim
- OAuth & Permissions -> Bot Token Scopes
  - chat:write
  - chat:write.public
  - channels:read
  - im:read
  - im:write
  - mpim:read
  - mpim:write
  - users:read
- Install to workspace -> Allow

&nbsp;

---

&nbsp;

## 05-slack-bot-uploads-file

- `.env`:
  - `SLACK_BOT_TOKEN_2=`
  - `SLACK_SOCKET_TOKEN_2=`
- [slack api](https://api.slack.com/)
- Create slack workspace -> Create app
- Connect using Socket Mode -> Generate socket-token
- Event Subscription -> Enable Events
- OAuth & Permissions -> Bot Token Scopes
  - channels:read
  - chat:write
  - files:read
  - files:write
  - im:read
  - im:write
  - mpim:history
  - remote_files:read
  - remote_files:share
  - remote_files:write
- Install to workspace -> Allow

&nbsp;

---

&nbsp;

## 06-email-verifier-tool

&nbsp;

---

&nbsp;

## 07-simple-aws-lambda-example

- [Amazon - Lambda execution role](https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html)
- [Amazon - region=ap-southeast-1#/roles](https://us-east-1.console.aws.amazon.com/iamv2/home?region=ap-southeast-1#/roles)

```sh
aws iam create-role --role-name lambda-ex --assume-role-policy-document '{"Version": "2012-10-17","Statement": [{ "Effect": "Allow", "Principal": {"Service": "lambda.amazonaws.com"}, "Action": "sts:AssumeRole"}]}'

# Already set with previous command
aws iam create-role --role-name lambda-ex --assume-role-policy-document file://07-simple-aws-lambda-example/trust-policy.json

aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

cd 07-simple-aws-lambda-example
go build main.go
zip function.zip main
aws lambda create-function --function-name go-simple-aws-lambda-example --zip-file fileb://function.zip --handler main --runtime go1.x --role arn:aws:iam::560476749134:role/lambda-ex

aws lambda invoke --function-name go-simple-aws-lambda-example --cli-binary-format raw-in-base64-out --payload '{"What is your name?": "Jim","How old are you?": 33}' output.txt
```

&nbsp;

---

&nbsp;
