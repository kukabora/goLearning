module task

go 1.19

replace (
	sample.com/wordz => ../wordz
	taskModule => ../taskModule
)

require (
	github.com/huandu/xstrings v1.4.0 // indirect
	sample.com/wordz v0.0.0-00010101000000-000000000000 // indirect
	taskModule v0.0.0-00010101000000-000000000000 // indirect
)
