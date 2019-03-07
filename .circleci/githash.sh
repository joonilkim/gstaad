#!/bin/bash

cur=`git log --pretty=format:'%H' -n 1`
target=`git log --pretty=format:'%H' -n 1 -- $1`

if [ $cur = $target ] ; then true ; else false ; fi
