module worker-monitor-client

go 1.14

replace local.packages/timer => ./timer

replace local.packages/monitor => ./monitor

replace local.packages/cmd => ./cmd

replace local.packages/clientdata => ./clientdata

replace local.packages/scheduler => ./scheduler

replace local.packages/configloader => ./configloader

require (
	local.packages/configloader v0.0.0-00010101000000-000000000000
	local.packages/monitor v0.0.0-00010101000000-000000000000
	local.packages/scheduler v0.0.0-00010101000000-000000000000
	local.packages/timer v0.0.0-00010101000000-000000000000
	local.packages/clientdata v0.0.0-00010101000000-000000000000
)
