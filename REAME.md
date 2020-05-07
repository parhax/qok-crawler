
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