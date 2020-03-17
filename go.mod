module github.com/zoli/lime-tcell

go 1.13

require (
	github.com/gdamore/tcell v1.3.0
	github.com/limetext/backend v0.0.0-20200304080108-aaa3040e8347
	github.com/limetext/commands v0.0.0-20200315181857-449329163b0b
	github.com/limetext/text v0.0.0-20200304072429-a501ee418129
	github.com/limetext/util v0.0.0-20160325174435-20e1a4a3505f
	github.com/zoli/sublime v0.0.0-20200317202733-39f5b6be4e77
)

replace github.com/limetext/backend => ../backend

replace github.com/limetext/commands => ../commands

replace github.com/zoli/sublime => ../sublime
