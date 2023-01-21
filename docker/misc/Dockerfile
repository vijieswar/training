FROM ubuntu

RUN apt update 

RUN apt install git vim -y

RUN mkdir /code
RUN mkdir /root/.ssh

WORKDIR /code

RUN echo "source /code/config/git.sh" >> ~/.bashrc

CMD bash
