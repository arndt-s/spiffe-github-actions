import * as core from '@actions/core';
import { init } from './init';
import * as grpc from '@grpc/grpc-js';

/**
 * The main function for the action.
 *
 * @returns Resolves when the action is complete.
 */
export async function run(): Promise<void> {
    try {
        const idToken = await core.getIDToken();
        const workloadEndpoint = core.getInput('spiffe-workload-endpoint');
        
        const spiffeId = await initSidecar(workloadEndpoint, idToken);
      
        core.setOutput('spiffeid', spiffeId);
      } catch (error) {
        core.setFailed(error.message);
      }
  }

export async function initSidecar(workloadEndpoint: string, idToken: string): Promise<string> {
    const client = new init.InitAPIClient(workloadEndpoint, grpc.credentials.createInsecure());
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
    return response.spiffe_id;
}