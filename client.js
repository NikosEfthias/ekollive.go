const net = require("net")
var sock= new net.socket()
sock.connect({port:1111})
sock.on("data",function(d){console.log(d)})
setInterval(_=>{
	sock.write("pong")
},50)
