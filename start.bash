#!/bin/bash
echo "HelloIPFS"

GIT='git --git-dir='$PWD'/.git'

$GIT checkout hotfix
$GIT merge --no-ff master -m "Merged master"
$GIT push

start chrome "www.google.com";