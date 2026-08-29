# Goeasyenc
This is a very simple RC4 file encryptor/decryptor, kinda a mini ransomware, for an assignment I was given in a summer camp. Since I had a plan to learn the Go language before, I did this assignment in golang to pick up the language on the way. So ya, it's my first Go program xD.  
 
`rc4.go` encrypts and also decrypts `file.txt` if it runs twice. 
`eww.go` encrypts and adds *0xdeadbeef* at the end of the file. this is to differentiate if the file has been encrypted or not. because if it's already encrypted, running the encryptor decrypts it back. this is the behavior of rc4 actually. That's why I implemented that *0xdeadbeef* mechanism.  
`yummy.go` decrypts the file and has the *0xdeadbeef* mechanism too to differentiate if it's been decrypted. 

