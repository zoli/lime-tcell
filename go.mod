module github.com/zoli/lime-tcell

go 1.13

require (
	github.com/gdamore/tcell v1.3.0
	github.com/limetext/backend v0.0.0-20191206170531-4aa255549774
	github.com/limetext/commands v0.0.0-20191121171555-02d601171b26
	github.com/limetext/text v0.0.0-20200304072429-a501ee418129
	github.com/limetext/util v0.0.0-20160325174435-20e1a4a3505f
)

replace github.com/limetext/commands => ../commands

replace github.com/limetext/backend => ../backend
