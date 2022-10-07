# FGA Kominfo Learning

## Day 9

Authentication:
    - apa itu authentication
        -> mengidentifikasi bahw request itu diminta oleh user atau client yang authentic dan dikenali oleh server.
            masukin username dan password.

    - kenapa ada authentication
        -> supaya request dikenali oleh server. kalau misal ga dikenali giman? 
            - mengembalikan error ke client
                - USER NOT FOUND
                - WRONG PASSWORD
        
    - apa bedanya dengan authorization
        apakah kita berhak merequest/mengunjungi halaman atau API yang kita HIT. Kalaupun berhak, apa aja yang boleh dilakukan.
            - UNAUTHORIZED

    - jenis method authentication:
    authentication / authorization akan diberikan client ke server dalam bentuh PAYLOAD HEADER (key: Authorization)
        - Basic Auth: langsung memberikan username dan password, dalam bentuk STRING ENCODED BASE64
        - OAuth: memberikan informasi dalam bentuk string TOKENIZE
            - JWT (JSON Web Token)

    - real case condition and security

---