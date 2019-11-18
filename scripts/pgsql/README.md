```
cd $GOPATH/src/github.com/open-falcon/falcon-plus/
psql -U postgres -h 127.0.0.1 < scripts/pgsql/db_schema/0_create_user.sql
psql -U postgres -h 127.0.0.1 < scripts/pgsql/db_schema/1_uic-db-schema.sql
psql -U postgres -h 127.0.0.1 < scripts/pgsql/db_schema/2_portal-db-schema.sql
psql -U postgres -h 127.0.0.1 < scripts/pgsql/db_schema/3_dashboard-db-schema.sql
psql -U postgres -h 127.0.0.1 < scripts/pgsql/db_schema/4_graph-db-schema.sql
psql -U postgres -h 127.0.0.1 < scripts/pgsql/db_schema/5_alarms-db-schema.sql
```
