[bristleurl]: https://bristle.io
[githuburl]: https://github.com/bristleio/docker-backup
[hub]: https://hub.docker.com/r/bristleio/backup/

[bristle.io][bristleurl]
======

The [bristle.io][bristleurl] team has created a container to do docker volume backups

bristleio/docker-backup
======
[![Version](https://img.shields.io/github/tag/bristleio/docker-backup.svg)][githuburl][![Docker Pulls](https://img.shields.io/docker/pulls/bristleio/backup.svg)][hub][![Docker Stars](https://img.shields.io/docker/stars/bristleio/backup.svg)][hub][![Build Automated](https://img.shields.io/docker/automated/bristleio/backup.svg)][hub][![Build Status](https://img.shields.io/docker/build/bristleio/backup.svg)][hub]


Usage
------

```sh
docker run \
  --rm \
  -v application_config:/src \
  -v application_backup:/dest \
  bristleio/backup
```

Parameters
------

The parameters that are used to facilitate the backup are the volumes that attach to `/src` and `/dest`. `/src` will always copy to `/dest` but it will not remove items on `/dest`. The volume attached to `/dest` should be a new volume.

Validation
-----

This can be run to validate the volume backup.

```sh
docker run \
  -it \
  -v application_backup:/backup \
  busybox /bin/sh
```

Once in the container, check the contents of `/backup`.

Please reach out to us if you see issues.
