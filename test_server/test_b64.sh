cat testCiphers | go run ../... -u "http://localhost:5000/decrypt?cipher=$" -err IncorrectPadding -p 200 -post "cipher=$" -cookie "token=$" "$@" > out
