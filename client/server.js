const grpc = require("grpc")

const messages = require("./proto/service_pb")
const services = require("./proto/service_grpc_pb")

function main() {
  const client = new services.AddServiceClient("localhost:4040", grpc.credentials.createInsecure())

  let request = new messages.Request()

  request.setA(2)
  request.setB(99)

  client.multiply(request, function(err, response) {
    console.log(response)
  })
}

main()