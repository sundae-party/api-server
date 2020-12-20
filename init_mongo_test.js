db = db.getSiblingDB($DB_NAME)
db.createUser(
    {
        user: $DB_USER,
        pwd: $DB_PASS,
        roles: [
            {
                role: "readWrite",
                db: $DB_NAME
            }
        ]
    }
);