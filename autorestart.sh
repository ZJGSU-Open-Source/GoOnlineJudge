#! /bin/sh
while true
do
pid=`pgrep GoOnlineJudge`
if [ $? -ne 0 ]
then
    echo "not running"
    echo "At time:`date`: Oj is down.Restarted successful." >> oj_status_log
    ./GoOnlineJudge &
else
    echo "is running"
fi
done
