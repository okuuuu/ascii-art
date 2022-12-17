# Test script

go run main.go --align right something standard
sleep 2

go run main.go --align=right left standard
sleep 2

go run main.go --align=left right standard
sleep 2

go run main.go --align=center hello shadow
sleep 2

go run main.go --align=justify "1 Two 4" shadow
sleep 2

go run main.go --align=right 23/32 standard
sleep 2

go run main.go --align=right ABCabc123 thinkertoy
sleep 2

go run main.go --align=center "#$%&\"" thinkertoy
sleep 2

go run main.go --align=left '23Hello World!' standard
sleep 2

go run main.go --align=justify 'HELLO there HOW are YOU?!' thinkertoy
sleep 2

go run main.go --align=right "a -> A b -> B c -> C" shadow
sleep 2

go run main.go --align=right abcd shadow
sleep 2

go run main.go --align=center ola standard
sleep 2
