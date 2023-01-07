FROM rockylinux:9

RUN useradd -ms /bin/bash go-user

USER go-user

WORKDIR /home/go-user

ENTRYPOINT ["executable", "param1", "param2"]

CMD ["-c"]