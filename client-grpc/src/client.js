const {EchoRequest, EchoResponse} = require('./echo_pb.js');
const {EchoClient} = require('./echo_grpc_web_pb.js');

var echoService = new EchoClient('http://localhost');

var request = new EchoRequest();
request.setContent('Hello Echo!');

echoService.echo(request, {}, function(err, response) {
  if (err) {
    console.log(err.code);
    console.log(err.message);
  } else {
    console.log(response.getContent());
    document.getElementById("grpc_resp").innerHTML = response;
  }
});
