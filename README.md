# Monitoring Security Issue in ORG or Github Project



* How To Setup

```sh
$ export git_notif=<your github token> 
$ export owner=<your org or owner repo>
```


* How to user
```sh
gitnotif-security ~% gitnotif  -h 
Usage of gitnotif:
  -repo string
    	for search in repo (default "false")
  -verbose
    	Verbose mode
```

* Scan All Repo 

```sh
$ gitnotif --verbose=true
```

* Scan Single Repo
```sh
$ gitnotif --repo=my-repo
```
