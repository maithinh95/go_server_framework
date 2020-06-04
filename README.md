# README #

START WITH:
$ make

FOR DOCKER:
- Build: $ docker build -t imagename:tag -f Dockerfile.tagfile .
- Run:   $ docker run -it --rm imagename:tag
         $ docker run -d -v share_dir:SHARE_DIR -v log_dir:LOG_DIR -e SHARE_DIR=share_dir -e LOG_DIR=log_dir --name dockername --restart on-failure imagename:tag 