#! /bin/sh
# use `crontab -e` to add the script so that it can backup automatically
# e.g 
# 0 5 * * * /home/acm/go/src/GoOnlineJudge/maintain/db_backup.sh
# means run the script at 5 o'clock every morning
mongodump -d oj -o ../../db_backup
