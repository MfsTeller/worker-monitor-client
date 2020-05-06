module monitor

go 1.14

replace local.packages/timer => ../timer

replace local.packages/cmd => ../cmd

require (
	local.packages/cmd v0.0.0-00010101000000-000000000000
	local.packages/timer v0.0.0-00010101000000-000000000000
)
