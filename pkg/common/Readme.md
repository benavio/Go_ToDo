RUN db
docker run -it --name GormPostgresDB -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=user -e POSTGRES_DB=db -p 5432:5432 postgres

view db in docker
docker exec -it <<containerName>> /bin/bash 
psql -U user -d db \dt