# Spinnaker Demo

This is just a quick demo application to show off the abilities of Spinnaker and deploying to Google Cloud.

The app will print out the Google availability zone and a version number to show that the app was updated on deployments
and to demo Spinnaker rollbacks. 

Very basic Go app using net/http for request handling.

# Building/Releasing

This project is setup to use Travis CI to build on a change to master. Travis CI usines FPM to generate
a deb package with the Go binary and init.d script, versioned based on the Travis job number. It uploads this deb to PackageCloud.

Spinnaker is setup to monitor the Travis job for successful compeletion and then install the proper version. 


