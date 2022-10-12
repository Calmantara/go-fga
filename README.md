# FGA Kominfo Learning

## Day 10

Context:
    - apa itu context
    context is a standard package of Golang that makes it easy to pass request-scoped values, cancelation signals, and deadlines across API boundaries to all the goroutines involved in handling a request.
    
    - context di golang
    - use case:
        - deadline & timeout memberikan deadline pada context dan mengecheck apakah context masih berada dalam range waktu tertente
        - cancel ketika ctx sudah tidak valid lagi, biasana cancel ini akan selalu dipanggil di akhir function (bp: defer)
        - value kita bisa memberikan key-value pair, dan mengakses key-value pair tsb pada method/function manapun yang menerima suatu context

Middleware:
    - apa itu middleware
    - kenapa middleware ada
    - middleware x cotext

    middleware bisa diassign di:
        - gin router group
        - gin method (POST, GET, etc)
        - gin.USE di paling depan (berlaku untuk semua request yang masokk)
JWT:
    - encoding and decoding concept
        - base64
        - hash
    - apa itu jwt
    - kenapa jwt ada
    - bagian dari JWT
        - header
        - body
        - signature
    - playing around with katara's jwt lib
        - encode
        - decode
        - verify
https://www.base64decode.org/
https://emn178.github.io/online-tools/sha1.html
https://jwt.io/
---
[!alt_day10](./day10.drawio.png)