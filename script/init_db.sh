#!/bin/sh

sudo su - postgres -c "createuser fib_app"
sudo su - postgres -c "createdb fib_db"
sudo -u postgres psql -U postgres -d fib_db -c "GRANT ALL PRIVILEGES ON DATABASE fib_db TO fib_app"