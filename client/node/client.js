const demo_proto = require('./proto')
const grpc = require("@grpc/grpc-js")

function main() {
  let client = new demo_proto.Demo('127.0.0.1:50054', grpc.credentials.createInsecure())
  client.Add({ x: 10, y: 2 }, function(err, response) {
    if (err) {
      console.error('Error: ', err)
    } else {
      console.log(response.result)
    }
  })
}

main()
