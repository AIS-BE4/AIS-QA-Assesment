Run ```docker compose up```<br />
This run should end up in an error, as database has not been initiated<br />
Enter store_service container with ```docker exec -it <Container_ID> /bin/bash```<br />
Run ```npx sequelize db:create``` to create db<br />
Run ```npm i``` then Run ```npx sequelize db:migrate``` to migrate the tables<br />
Quit the container and re-run ```docker compose up```<br />
Please make sure that Port _5432_ is available for PgSQL and port _8000_, _8001_, _8002_, _8003_, _8004_ and _8005_ are available
