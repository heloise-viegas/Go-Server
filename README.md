# Go-Server
A simple web server written in Go

http package

1) handle(string,http.handler)
handler can be http.fileserver(http.dir("./folderName")
2) handlefunc(string,func)   
where func = func(w response writer,r http.request)
3)listenandserve(string addr,handler)
listens to address and then calls serve with the handler to handle the request
handler is set to nil
this func mostly returns nil
4)http.Error(w http.ResponseWriter, error string, code int)
