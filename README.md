# govtech_onecv_assessment
Name: Lim Boon Hai

This project is done as part of the submission for GovTech OneCV's Golang take home assessment.
This project is done using Gin and GORM.

To run the project, please follow the steps below:
1. Navigate to your desired directory and clone the repository using `git clone https://github.com/boonhaii/govtech_onecv_assessment.git`
2. Navigate to the cloned directory.
3. Start the program via the command `go run main.go`. The API should start up on Port 8080. Please also do ensure that you have started your MySQL server, and have the database with the required tables present. For production, the database name is `apidb`, and for test, the database name is `apidb_test`. Both databases have 3 tables: `Teachers`, `Students` and `Registers`. Please do refer to `Config/db_setup.sql` for the commands to create the database.
4. Visit `localhost:8080` to verify that the API is working. You should see `Welcome to this API` being shown on the browser.
