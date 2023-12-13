# Developer comments

This application was build based on the requirements written in the ASSIGNMENT.md file together with my interpretation of existing code and some creative logic. I tried not to take so much time on this test project but still my whish was to look and feel clean while person who is going to review it does not have difficult time.

# How to run server

Please execute run_server.sh files to start server. Server will listen for all client requests. There is no additional configuration for the server.

# How to run client

To run a client and test the application run go run ./test/client/main.go and please take a look at ./app/client/main.go file where StartClient() function is called. There you can find examples on how to run app for default "Gothenburg" tax data and also there is way to simulate bonus scenario if you uncomment //SendGetRequest("Beograd") line.

# There are some tests

While i was writing and testing the app i managed to write few tests that i find use of and they can be executed by running script run_tests.sh. 