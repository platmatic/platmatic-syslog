var syslog = require("syslog-client");

var options = {
    transport: syslog.Transport.Tcp,
    port: 8080
};

var client = syslog.createClient("127.0.0.1", options);

client.log("example message");