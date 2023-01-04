module application

go 1.19

replace structs/customer => ../Customer

require structs/customer v0.0.0-00010101000000-000000000000 // indirect
