# FGA Kominfo Learning

## Day 11

- Testing in Development
     - SDLC (software development life cycle)
        sdlc -> life cycle mulai dari menerima req, planning, designing, implementation, test, hingga evaluation
     - STLC (software test life cycle)
        stlc -> test life cycle (yang biasa dilakukan dev / qa) untuk membuat test case hingga automation test
     - Testing in software development
        - unit test
            test yang dibuat untuk mengetest semua scenario yang bisa terjadi di suatu function (dev)
        - integration test
            test yang dilakukan untuk mencoba integrasi dalama 1 fitur (baik cross team atau dalam 1 code base) (dev) 
        - qa test
            test yang dilakukan baik secara manual atau otomatis untuk menjawab test case yang sudah dibuat dalam stlc (qa)
            - automation test
            - manual test
            - security test (security engineer)
            - load test
                supaya kita tau apakah logic codingan kita, sudah optimized atau belum
                sehingga bisa menampung load / request dalam waktu yang bersamaan sangat banyak
        - UAT test 
            test yang dilakukan oleh pemberi req (business team), untuk memastikan tidak adanya missed req, tidak ada bug, dan siap untuk naik production (business)
        - prod test / smoke test
            test singkat yang dilakukan ketika sudah naik ke production (qa)
    
- golang:
   - unit test
   - test coverage 
       mendeteksi sudah berapa banyak unit test yang tercover untuk setiap method/function di suatu workspace
       Test coverage sendiri dibagi menjadi 2:
           - apakah semua function sudah ada unit testnya?
           - apakah semua conditional statement (if else, switch case) dalam satu function udah tercover semua atau belum

- Code quality (sonarqube):
   - code coverage
   - code duplication
     untuk mencapai suatu close loop test, kita bisa mocking data dengan menggunakan gomock
     https://www.youtube.com/watch?v=KJXXboJz7BA&t=849s&ab_channel=AminMir

example repo for gomock:
git@github.com:Calmantara/go-playground.git

untuk menjalankan test di go
1. menggunakan debugger test (vscode, goland)
2. go test ./...
3. go test 
4. go test -coverprofile cov.out ./...
