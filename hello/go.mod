module example/hello

go 1.18

replace example/first => ../first

replace example/greetings => ../greetings

require example/greetings v0.0.0-00010101000000-000000000000
