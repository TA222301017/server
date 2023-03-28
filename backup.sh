echo "Loading .env file"
for var in $(grep -v '^#' .env | xargs)
do
    export $(echo $var | tr -d \n)
done

sudo -u postgres pg_dump ${DB_NAME//[$'\t\r\n ']} > ${DB_NAME//[$'\t\r\n ']}.sql