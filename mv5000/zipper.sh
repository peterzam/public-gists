#!/bin/bash
# Used : tar, pigz, split

opt=$1
dir=$2
size=2GB

if [[ $opt == '-c' ]]
then
        tar -c "${dir}" | pigz -p 4 -c | split -a 2 --additional-suffix=".tar.gz" -d -b "${size}" - "${dir}."
elif [[ $opt == '-d' ]]
then
        cat $dir* | pigz -p 4 -dc | tar xf -
else
        echo "zipper.sh <option> <dir>"
        echo "option : -d decompress"
        echo "       : -c compress"
fi