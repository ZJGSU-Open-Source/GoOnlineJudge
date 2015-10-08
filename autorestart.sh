#! /bin/sh

while true
do
pid=`history | grep ojapi`
if [ $? -ne 0 ]
then
    echo "Not running"
    echo "At time:`date`: OJ is down. Restart succeeded." >> oj_status_log
    restweb run ojapi &
else
    echo "At time:`date`: is running"
fi
done
