FROM python:3.7-slim

ADD ./ app/

# Create app directory
WORKDIR /app

# Install app dependencies
COPY requirements.txt ./

ENV PYTHONPATH "${PYTHONPATH}:/app/scedules/:/app/models/:/app/processing/:/app/helpers/"

RUN pip install -r requirements.txt

# Bundle app source
COPY . /app

EXPOSE 2002

CMD [ "python", "main.py" ]