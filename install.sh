#!/usr/bin/bash

# Load .env file
echo "Loading .env file"
for var in $(grep -v '^#' .env | xargs)
do
    export $(echo $var | tr -d \n)
done
export PGPASSWORD=${DB_PASSWORD//[$'\t\r\n ']}

DEPS="wget unzip ca-certificates"
SERVER_REPO=ta222301017/server
ADMIN_PANEL_REPO=ta222301017/admin-panel-js

CWD=$(pwd)/${APP_NAME//[$'\t\r\n ']}
USERNAME=$(id -un)
ARCH=$(uname -m)

echo "Fething latest tags..."
SERVER_TAG=$(curl --silent "https://api.github.com/repos/$SERVER_REPO/releases/latest" | 
    grep '"tag_name":' |                                            
    sed -E 's/.*"([^"]+)".*/\1/')

ADMIN_PANEL_TAG=$(curl --silent "https://api.github.com/repos/$ADMIN_PANEL_REPO/releases/latest" | 
    grep '"tag_name":' |                                            
    sed -E 's/.*"([^"]+)".*/\1/')


echo "Determining machine architecture..."
if [ $ARCH == "x86_64" ] 
then
    ARCH=amd64
elif [ $ARCH == "i386" ] 
then
    ARCH="386"
else
    ARCH=arm
fi

echo "Installing $DEPS"
apt -qq install $DEPS --silent -y

echo "Installing postgresql..."
wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -
sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt/ $(lsb_release -cs)-pgdg main" >> /etc/apt/sources.list.d/pgdg.list'
apt -qq update
apt -qq install postgresql postgresql-contrib -y

echo "Creating postgresql user and database..."
sudo -u postgres psql -w -c "CREATE USER ${DB_USER//[$'\t\r\n ']} WITH PASSWORD '${DB_PASSWORD//[$'\t\r\n ']}' CREATEDB;"
createdb -h localhost -p 5432 -U ${DB_USER//[$'\t\r\n ']} -w ${DB_NAME//[$'\t\r\n ']}

mkdir -p $CWD
cp .env $CWD/.env

echo "Downloading latest server executable..."
wget --quiet https://github.com/TA222301017/server/releases/download/$SERVER_TAG/server-linux-$ARCH -O $CWD/server
chmod +x $CWD/server

echo "Downloading latest admin panel build..."
wget --quiet https://github.com/TA222301017/admin-panel-js/releases/download/$ADMIN_PANEL_TAG/dist.zip -O dist.zip
unzip dist.zip
mv dist $CWD/static
rm dist.zip

echo "Creating daemon..."
cat > $CWD/${APP_NAME//[$'\t\r\n ']}.service << EOF

[Unit]
Description=Service for $APP_NAME
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
User=$USERNAME
WorkingDirectory=$CWD
ExecStart=$CWD/server
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target

EOF

link $CWD/${APP_NAME//[$'\t\r\n ']}.service /etc/systemd/system/${APP_NAME//[$'\t\r\n ']}.service
systemctl daemon-reload
systemctl start ${APP_NAME//[$'\t\r\n ']}.service

echo "Done, server started on $APP_HOST:$APP_PORT"