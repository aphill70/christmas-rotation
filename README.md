# christmas-rotation

[![Build Status](https://travis-ci.com/aphill70/christmas-rotation.svg?branch=master)](https://travis-ci.com/aphill70/christmas-rotation)

Simple app written to store and generate a christmas present rotation using Google Sheets.

#TODO
- Handle case when some but not all participants rollover
- Implement CLI interface so sheet and other variables aren't hardcoded
- fix sheets_test.go: Interface change never got fixed
- Rollover Needs to be more idempotent probably need to use an array to guarantee accurate rollover and ordering probably need a year to array index mapping and need to calculate the rollover year independent of actually calculating the eligible givers.