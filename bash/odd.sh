X=1
while [ $X -le 99 ]
do
 rem=$(($X % 2))
 if [ "$rem" -ne "0" ];then
    echo $X
 fi
 X=$((X+1))
done 