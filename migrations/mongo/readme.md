```shell

use mongovspostgres

db.states.insert({name: "SC", country: "Brasil"})

db.cities.insert({name: "Joinville", state_id:ObjectId("647798d3939fc4752dc8d081")})

db.users.insert({name: "Leonardo Miranda", city_id: ObjectI("647798d3939fc4752dc8d081")})

```