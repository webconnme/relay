#!/bin/sh
TARGET_FILES="app_relay example/console"
for FILE in ${TARGET_FILES}
do
  go install ${FILE}
done
