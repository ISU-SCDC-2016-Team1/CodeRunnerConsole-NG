# CodeRunnerConsole-NG

CodeRunnerConsole-NG is the next generation of code runner interfaces,
developed here at the Cluster Deployment Company. Please make sure your
public key is available on the runner through a keyescrow deployment.

##Usage:
	crconsole-ng -u <username> -i <key> -p <project> -r <runner> -x <redirect> <verb>

	Where username, key, project, and runner are all required flags, stdin is
	optional and verb is one of:
	  deploy, clean, build, run, stdin, get

	where stdin takes an additional argument of the stdin file, and get takes
	an additional argument of the io stream.

	Note that redirect must be specified as all in order to read stderr and
	stdout. Further, stdin is ignored on all commands except run. 
