rs.initiate({
    _id: "rs0",
    members: [
        {
            _id: 0,
            host: "mongo:27017",
        }
    ]
})

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