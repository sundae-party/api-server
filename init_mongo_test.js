db = db.getSiblingDB(process.en.DB_NAME)
db.createUser(
    {
        user: process.en.DB_USER,
        pwd: process.en.DB_PASS,
        roles: [
            {
                role: "readWrite",
                db: process.en.DB_NAME
            }
        ]
    })