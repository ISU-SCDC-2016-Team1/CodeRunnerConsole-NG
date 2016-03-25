# CodeRunnerConsole-NG

CodeRunnerConsole-NG is the next generation of code runner interfaces,
developed here at the Cluster Deployment Company. Please make sure your
private key is available on the runner. 

##Usage:

	crconsole-ng -u <username> -i <key> -p <project> -r <runner> <verb>

	Where username, key, project, and runner are all required flags, stdin is
	optional and verb is one of:
	  deploy, clean, build, run, stdin, stdout

	where stdin takes an additional argument of the stdin file
