#!/bin/bash

mkdir $2;ls $1 | head -5000 | xargs -I{} mv ./$1/{} ./$2