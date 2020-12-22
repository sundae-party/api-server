rs.initiate()

sleep(2000)

db = db.getSiblingDB("sundae");
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
    });