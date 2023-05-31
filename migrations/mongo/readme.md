```shell

use mongovspostgres

db.states.insert({name: "SC", country: "Brasil"})

db.cities.insert({name: "Joinville", state_id:ObjectId("647798d3939fc4752dc8d081")})

db.users.insertMany([
    {name: "Leonardo Miranda 0", city_id: ObjectId("6477991c939fc4752dc8d082")},
    {name: "Leonardo Miranda 1", city_id: ObjectId("6477991c939fc4752dc8d082")},
    {name: "Leonardo Miranda 2", city_id: ObjectId("6477991c939fc4752dc8d082")},
    {name: "Leonardo Miranda 3", city_id: ObjectId("6477991c939fc4752dc8d082")},
    {name: "Leonardo Miranda 4", city_id: ObjectId("6477991c939fc4752dc8d082")},
    {name: "Leonardo Miranda 5", city_id: ObjectId("6477991c939fc4752dc8d082")},
    {name: "Leonardo Miranda 6", city_id: ObjectId("6477991c939fc4752dc8d082")},
    {name: "Leonardo Miranda 7", city_id: ObjectId("6477991c939fc4752dc8d082")},
    {name: "Leonardo Miranda 8", city_id: ObjectId("6477991c939fc4752dc8d082")},
    {name: "Leonardo Miranda 9", city_id: ObjectId("6477991c939fc4752dc8d082")},
])

```