db.createUser(
    {
        user: "sundae",
        pwd: "pass",
        roles: [
            {
                role: "readWrite",
                db: "sundae"
            }
        ]
    }
);