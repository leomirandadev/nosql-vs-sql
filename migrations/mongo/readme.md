```shell

use mongovspostgres

db.states.insertMany([
    {name: "SC", country: "Brasil"},
    {name: "SP", country: "Brasil"},
    {name: "PR", country: "Brasil"},
    {name: "RS", country: "Brasil"},
    {name: "ES", country: "Brasil"},
])

db.cities.insertMany([
    {
        name: "Joinville",
        state_id:ObjectId("6477c297245a0416d6a64bcb"),
    },
    {
        name: "Araquari",
        state_id:ObjectId("6477c297245a0416d6a64bcb"),
    },
    {
        name: "São Francisco do Sul",
        state_id:ObjectId("6477c297245a0416d6a64bcb"),
    },
    {
        name: "Blumenau",
        state_id:ObjectId("6477c297245a0416d6a64bcb"),
    },
    {
        name: "Itajaí",
        state_id:ObjectId("6477c297245a0416d6a64bcb"),
    },
    {
        name: "Brusque",
        state_id:ObjectId("6477c297245a0416d6a64bcb"),
    },
    {
        name: "Chapecó",
        state_id:ObjectId("6477c297245a0416d6a64bcb"),
    },
])

db.users.insertMany([
    {name: "Leonardo Miranda 0", city_id: ObjectId("6477c2e3c1adaf0e33063c02")},
    {name: "Leonardo Miranda 1", city_id: ObjectId("6477c2e3c1adaf0e33063c02")},
    {name: "Leonardo Miranda 2", city_id: ObjectId("6477c2e3c1adaf0e33063c02")},
    {name: "Leonardo Miranda 3", city_id: ObjectId("6477c2e3c1adaf0e33063c02")},
    {name: "Leonardo Miranda 4", city_id: ObjectId("6477c2e3c1adaf0e33063c02")},
    {name: "Leonardo Miranda 5", city_id: ObjectId("6477c2e3c1adaf0e33063c02")},
    {name: "Leonardo Miranda 6", city_id: ObjectId("6477c2e3c1adaf0e33063c02")},
    {name: "Leonardo Miranda 7", city_id: ObjectId("6477c2e3c1adaf0e33063c02")},
    {name: "Leonardo Miranda 8", city_id: ObjectId("6477c2e3c1adaf0e33063c02")},
    {name: "Leonardo Miranda 9", city_id: ObjectId("6477c2e3c1adaf0e33063c02")},
])

```