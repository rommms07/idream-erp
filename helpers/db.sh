#!/bin/bash -

echo "Creating database '$MYSQL_DB_NAME'..."
mysql -u$MYSQL_USER -p$MYSQL_PASSWORD -e "create database $MYSQL_DB_NAME;"
