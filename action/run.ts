import * as core from '@actions/core';
import * as grpc from '@grpc/grpc-js';
import { InitAPIClient, InitRequest, InitResponse } from './init';
import { JWTSVID, JWTSVIDRequest, JWTSVIDResponse, SpiffeWorkloadAPIClient } from './workload';

const defaultSpiffeWorkloadEndpoint = 'unix:/spiffe/workload.sock';

/**
 * The main function for the action.
 *
 * @returns Resolves when the action is complete.
 */
export async function run(): Promise<void> {
    try {
        const action = core.getInput('action');
        
        if (action == "init") {
            await initAction();
        } else if (action == "fetch-jwt-svid") {
          await fetchJWTSVIDAction();
        } else {
            core.setFailed(`Invalid action: ${action}`);
        }
        
      } catch (error) {
        core.setFailed(error.message);
      }
  }


async function initAction(): Promise<any> {
    const idToken = await core.getIDToken();

    const client = new InitAPIClient(getSpiffeWorkloadEndpoint(), grpc.credentials.createInsecure());
    const initRequest = new InitRequest({ id_token: idToken });
    const response = await new Promise<InitResponse>((resolve, reject) => {
        client.Init(initRequest, (err: any, response: InitResponse | undefined) => {
          if (err || !response) {
            reject(err);
          } else {
            resolve(response);
          }
        });
      });
    core.setOutput('spiffe_id', response.spiffe_id);
}

async function fetchJWTSVIDAction(): Promise<any> {
  const audienceParam = core.getInput('audience')
  const spiffeID = core.getInput('spiffe_id');

  var audiences = [];
  if (audienceParam) {
    audiences = audienceParam.split(",");
  }

  const client = new SpiffeWorkloadAPIClient(getSpiffeWorkloadEndpoint(), grpc.credentials.createInsecure());
  const response = await new Promise<JWTSVIDResponse>((resolve, reject) => {
      const jwtSvidRequest = new JWTSVIDRequest({ audience: audiences, spiffe_id: spiffeID });
      client.FetchJWTSVID(jwtSvidRequest, (err: any, response: JWTSVIDResponse | undefined) => {
        if (err || !response) {
          reject(err);
        } else {
          resolve(response);
        }
      });
    });
  
  if (response.svids.length > 1) {
    core.setFailed('Expected a single JWT-SVID, but received multiple.');
  } else if (response.svids.length == 0) {
    core.setFailed('Expected a single JWT-SVID, but received none.');
  }

  const svid = response.svids[0];
  core.setSecret(svid.svid);
  core.setOutput('jwt_svid', svid.svid);
  core.setOutput('spiffe_id', svid.spiffe_id);
}

function getSpiffeWorkloadEndpoint(): string {
  return core.getInput('spiffe_workload_endpoint') ?? defaultSpiffeWorkloadEndpoint;
}