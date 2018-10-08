FROM python:3-alpine
LABEL maintainer="@douglaspands"
LABEL version="1.0.0"
COPY . /var/www
WORKDIR /var/www
RUN pip install -r requirements.txt
CMD python app.py