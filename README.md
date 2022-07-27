
# AIS QA ASSESMENT


## Setup Database

Make Sure you already install docker first before you doing this test

Go to folder ```pg-docker``` using this command: ```cd pg-docker```

Then Run ```docker-compose up -d```

NOTE:

```To Check the database you can use any database client(such as pgAdmin / Navicat), then connect to host: localhost || username: postgres || password: postgres || port: 5435 || initial datbase: postgres```


## Preparation: Running Service

Make Sure you already install docker first before you doing this test

Run ```docker compose up```

This run should end up in an error, as database has not been initiated


Enter store_service container with ```docker exec -it <Container_ID> /bin/bash```

(```Example: docker exec -it e90b8831a4b8 /bin/bash```)

```(For Getting Docker Container ID, you can run `docker ps` then find container named store_service)```

Run ```npx sequelize db:create``` to create db

Run ```npm i``` then Run ```npx sequelize db:migrate``` to migrate the tables

Quit the container and re-run ```docker compose up```

Please make sure that Port _5432_ is available for PgSQL and port _8000_, _8001_, _8002_, _8003_, _8004_ and _8005_ are available



## Question

After all the services running properly, give us some PoW(Proof of Work) / PoC(Proof of Concept) The way you running automation test on this services.
Please make a list of all the scenario that you tested on the service.
