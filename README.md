# spiffe-github-actions

Proof of concept to offer a [SPIFFE Workload Endpoint](https://github.com/spiffe/spiffe/blob/main/standards/SPIFFE_Workload_Endpoint.md) on Github Actions.

If the internet has led you here and you are interested in using SPIFFE in Github Actions feel free to contact me at arndts.ietf@gmail.com.

## Scope

* Code focuses on the SPIFFE part only. 
* Code does NOT validate the Github ID Token.
* Code only issues a dummy JWT-SVID.

## Limitations

* Github service containers don't have access to the ID token. The `init` step & corresponding gRPC API is used to initialize the service container with an ID token obtained from the main job.
* Will only work in Linux. Did not test MacOS runners.