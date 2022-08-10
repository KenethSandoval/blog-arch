import { LoginRequest } from "@proto/api/protobuf/auth_pb";
import { AuthServiceClient } from "@proto/api/protobuf/auth_grpc_web_pb";


export const client = new AuthServiceClient("http://localhost:50051", null, null);

export const login = async () => {
	let loginRequest = new LoginRequest({});

	let rest = await doCall(loginRequest); 
	console.log(rest);
}

function doCall(loginRequest: any) {
  return new Promise((resolve, reject) => {
        client.login(loginRequest, null, (err, response) => {
            if (err) {
                reject(err.message);
            }
            else {
                resolve(response);
            }
        });
   });
}
