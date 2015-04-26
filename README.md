mini-go-cluster
==================
Mini cluster to execute parallel jobs using go lang.
Jobs can be either:
1. HTTP/ API requests
2. Terminal commands

Currently supports HTTP GET requests only.


## You Will Have
An output file containing jobs results.
Currenly it supports only web requets.


## Usage
Execute go cluster by running the command:
<br/><code>
go run src/bootstrap.go <input file path> <output file path> <error file path> <log file path>
</code>

Example:
<br/><code>
go run src/bootstrap.go /www/input/input.txt /www/output/output.txt /www/output/output_error.txt
</code>


## Configuration
You can modify the following config for number of jobs to execute parallel
<br/><code>
var MAX_CONCURRENT_CONNECTION = 10
</code>


## Future Development
As mentioned currently it supports HTTP GET requests. Later it is going to support POST, PUT and DELETE requests.
Terminal commands will also be supported by system.


## Developer Resources
Check out the URLs bellow to find out how its done:<br/>
[Go Lang Documentation and instalation](http://golang.org/)<br/>


## Interested in contributing?
If you wanna add more features and user options then just fork this repo from the link bellow:
https://github.com/waqar-alamgir/mini-go-cluster/fork


## Credits
mini-go-cluster by [Waqar Alamgir](http://waqaralamgir.tk)<br/>
[Web](http://waqaralamgir.tk)<br/>
[Twitter](http://www.twitter.com/wajrcs)
