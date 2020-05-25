# Iteration Names

### This Cloud Function generates an iteration name based on input.

#### How does it work?

This Cloud Function is deployed to a GCP project. See "How do I deploy the Cloud Function?" below for instructions.

Use the following curl command
`curl https://{zone-gcp-project}.cloudfunctions.net/GetIterationName`
with the suffix
`?letter=`
followed by a capital letter to generate an iteration name.

The Cloud Function uses the capital letter to randomly select an adjective and animal from map variables.
The input must be a single, capital letter, in a string.

The test checks that no unexpected errors occurred, and that the output begins with the correct letter.

#### Where does the Cloud Function live?

The Cloud Function is inside a project on Google Cloud Platform.

The curl command is invoking the `GetIterationName` function that was deployed there.
This GitHub repository allows me to make changes locally and redeploy if required.

#### How do I deploy the Cloud Function?

`gcloud functions deploy GetIterationName --runtime go113 --trigger-http`
 
Check you are in the correct region and project with gcloud locally.
The above command will deploy any changes made to the local code.
Once the changes have been deployed, you can check the code on [GCP](console.cloud.google.com).
Use the curl command to check that it functions as expected.
