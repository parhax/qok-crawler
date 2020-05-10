## Getting Started


For running this Go app you should first build its image by running:

` docker-compose up --build `

One HTTP endpoints have been provided in this app which you can test and examine it via Postman application or curl like following codes:

```
curl --location --request POST 'http://127.0.0.1:8686/crawl' \
--header 'Content-Type: text/plain' \
--data-raw '[
"https://quizofkings.com",
"http://google.com",
"http://aparat.com",
"http://martinfowler.com"
]
'
```