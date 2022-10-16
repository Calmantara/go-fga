# FGA Kominfo Learning

## Deployment

1. Docker - Docker Compose
```
    Mengurangi bentrok/ketergantunagn dependencies
    Term di Docker:
        - Container
        image yang running diatas docker orc
        - Image
        An image describes everything that is needed to create a container (ada binary golang app, ada suatu version tertentu)
        - Registry
        Storage untuk menyimpan image image kita
    
    Dockerfile
    file yang akan berisikan syntax untuk membuat image kita

    untuk membuat image, ada beberapa hal yang bisa dilakukan
    1. copy langsung code ke image golang (paling tidak recomended) kode bisa diakses ketika image dibuka
    2. build binary di local, terus copy binary ke base image alpine
    3. multi stage (copy code di golang image, lalu build binary di golang image, baru copy binary ke alpine image)
    https://docs.docker.com/engine/reference/builder/

    step by step membuat image:

    scenario1:
    paling tidak rekomended
        1. membuat dockerfile
        2. build image langsung
            docker build -t fga/userimage .
            (-t (tag) untuk memberikan nama image)
        3. untuk menjalankan image
            docker run -d --name fga-user1 -p 8080:8080 fga/user

    scenario2:
    biasanya binary go akan di buat di CI/CD environment (buat binary di CI/CD VM, lalu copy binary ke image dan membuat IMAGE)
        1. membuat dockerfile
        2. build go binary dengan define os dan arc
        GOOS=linux GOARC=amd64 go build  -o go-fga .
        3. build image langsung
            docker build -t fga/userimage .
            (-t (tag) untuk memberikan nama image)
        4. untuk menjalankan image
            docker run -d --name fga-user1 -p 8080:8080 fga/user
    
    scenario3:
    biasanya untuk mensupport CI/CD VM yang tidak ada GOLANG version. sehingga diperlukan GO Multistage Docker.
    Binary dibuat dalam satu serangkaian ketika membuat IMAGE
        1. membuat dockerfile
        2. build image langsung
            docker build -t fga/userimage .
            (-t (tag) untuk memberikan nama image)
        3. untuk menjalankan image
            docker run -d --name fga-user1 -p 8080:8080 fga/user

    Docker Compose
    Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application’s services.
```

2. SSH
3. VM/Server dan DNS (aws)
4. CI/CD in theory