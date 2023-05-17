# DNSlookUp
 The DNS LookUp tool is a command-line application written in Go 

Prerequisites

Go 1.16 or higher
Usage

To run the DNS LookUp tool, follow these steps:

Clone the repository:
git clone https://github.com/spix-777/DNSlookUp.git

Change into the project directory:
cd dns-lookup

Build the executable:
go
Copy code
go build

Run the tool:
./dns-lookup [options]
Replace [options] with the desired command-line options. The available options are:

-u (string): Specifies the search URL. Defaults to "example".
-f (bool): Specifies whether to display a message for failed DNS lookups. Defaults to false.
-t (bool): Specifies whether to display a message for successful DNS lookups. Defaults to false.
-a (bool): Specifies whether to display both failed and successful DNS lookups. Overrides -f and -t options. Defaults to false.

Example Usage

Perform DNS lookup for a specific URL:
./dns-lookup -u example.com
Display messages for both failed and successful DNS lookups:
./dns-lookup -u example.com -a
Output

The tool reads URLs from a file called output.txt and performs DNS lookups for each line. The results are displayed in the terminal based on the provided options.

Failed DNS lookups are shown in red, with the message "FAIL" appended to the URL.
Successful DNS lookups are shown in green, with the message "FOUND" appended to the URL.
Note: The colors may not be visible in all terminals.

License

This project is licensed under the MIT License. See the LICENSE file for details.