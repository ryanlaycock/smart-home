FROM tiangolo/uwsgi-nginx-flask:python3.7

WORKDIR /app

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

COPY . .

CMD ["python3", "-u", "server.py"]