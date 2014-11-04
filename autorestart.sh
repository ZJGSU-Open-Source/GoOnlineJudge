#! /bin/sh
status=`pgrep GoOnlineJudge`
if (status=='')
then
	echo "Restarting OJ"
	./GoOnlineJudge
fi
