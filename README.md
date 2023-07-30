## BTPN Syariah Fullstack Developer Virtual Internship Program

# User Endpoint :

    ✔️POST : /users/register
        - ID (primary key, required)
        - Username (required)
        - Email (unique & required) 
        - Password (required & minlength 6)
        - Relasi dengan model Photo (Gunakan constraint cascade)
        - Created At (timestamp)
        - Updated At (timestamp)
    ✔️GET: /users/login
        - Using email & password (required)
    ✔️PUT : /users/:userId (Update User)
    ✔️DELETE : /users/:userId (Delete User)

# Photos Endpoint

    ✔️POST : /photos 
        - ID
        - Title
        - Caption
        - PhotoUrl
        - UserID
    ✔️GET : /photos
    ✔️PUT : /photoId
    ✔️DELETE : /:photoId