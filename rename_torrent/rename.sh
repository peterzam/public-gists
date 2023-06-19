#!/bin/bash

for file in * ; do
    if [ -f "$file" ] ; then
        new_name=$(echo "$file" | sed -e 's/^.*] \s*//')
        mv "$file" "$new_name"
    fi
done

