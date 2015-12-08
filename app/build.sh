#!/bin/sh
TARGET_FILES="app_relay"
for FILE in ${TARGET_FILES}
do
  go install ${FILE}
done
