# podlog
Merge multiple remote pod io streams to local stdout.

Use case: You have some pods that don't share a common label or deployment (so the oc logs command is out) and you quickly want to know what they're up to.

# Requirements
- Authenticated openshift cli in $PATH
- All pods must log to stdio

# Install
Grab the binary for you os from the bin dir or build it yourself with the build script (requires go).

# usage
```bash
# Pass a list of resources like pods or deployments
$ podlog <pod1> <pod2> ...
```
Flags
- Use -version to print version and exit

Examples
```bash
# Stream logs from a subset of pods
$ podlog `oc get pods -o name | grep service`
# filter the combined output
$ podlog <pods> | grep <re>
```

# Todos
- [x] proto
- [x] enable grep
- [x] build script and install docs
- [ ] option to add pod name to output
- [ ] make -f optional
- [ ] publish image to some registry

# license
MIT
