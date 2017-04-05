FROM python:3.6.0-alpine
ADD . /src
WORKDIR /src
EXPOSE  8080
ENTRYPOINT ["python3", "-m", "http.server", "8080"]