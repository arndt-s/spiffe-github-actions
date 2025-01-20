import * as core from '@actions/core';
import { init } from './init';
import * as grpc from '@grpc/grpc-js';

try {
  const idToken = core.getInput('id-token');
  const client = new init.ProcessIDAttestationAPIClient("unix:/tmp/spiffe/workload.sock", grpc.credentials.createInsecure());

  const initRequest = new init.InitRequest({ id_token: idToken });
  const response = await new Promise<init.InitResponse>((resolve, reject) => {
    client.Init(initRequest, (err: any, response: init.InitResponse | undefined) => {
      if (err || !response) {
        reject(err);
      } else {
        resolve(response);
      }
    });
  });

  core.setOutput('spiffe_id', response.spiffe_id);
} catch (error) {
  core.setFailed(error.message);
}