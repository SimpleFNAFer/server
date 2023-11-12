# server
Web server for dos attack

## .env
```env
POSTGRES_USER: [pg_user]
POSTGRES_PASSWORD: [password]
POSTGRES_DB: [db]
POSTGRES_HOST: server-db

GF_SECURITY_ADMIN_USER: [gf_user]
GF_SECURITY_ADMIN_PASSWORD: [gf_password]
GF_DASHBOARDS_MIN_REFRESH_INTERVAL: 1s
```
****PUT .env FILE TO THE ROOT PROJECT DIRECTORY****
```
server
├── ...
├── docker-compose.yml
├── .env
└── ...
```
This is an example of .env file. \
It contains all variables that are mandatory for this project. \
Put ***your own values*** for variables having text in ***square brackets***. \
***Don't change*** `POSTGRES_HOST` and `GF_DASHBOARDS_MIN_REFRESH_INTERVAL` variables! \
\
(`POSTGRES_HOST` is a postgres container name. \
If you change it, go web server will not be able to connect to the database. \
`GF_DASHBOARDS_MIN_REFRESH_INTERVAL` sets the minimum refresh rate for dashboards. \
If you change it, you will not be able to set refresh rate to 1s, default is 5s)
