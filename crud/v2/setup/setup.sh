# createuser -P -d gwp  # password: gwp
# createdb gwp 
psql -U gwp -f setup.sql -d gwp