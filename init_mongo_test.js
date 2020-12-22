db = db.getSiblingDB("sundae")
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
    })

rs.initiate();