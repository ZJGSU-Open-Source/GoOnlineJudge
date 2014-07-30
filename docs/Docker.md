##Docker

```bash
sudo apt-get update
sudo apt-get install apt-transport-https
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 36A1D7869245C8950F966E92D8576A8BA88D21E9
sudo sh -c "echo deb https://get.docker.io/ubuntu docker main\
> /etc/apt/sources.list.d/docker.list"
sudo apt-get update
sudo apt-get install lxc-docker
```

To verify that everything has worked as expected:
`sudo docker run -i -t ubuntu /bin/bash`
Which should download the `ubuntu` image, and then start `bash` in a container.
