test:
	coverage run -m unittest discover test/ "*_test.py"

coverage:
	coverage report -m

run:
	python main.py

docker-build:
	docker build -t fetch-app -f Dockerfile .

docker-run:
	docker run -p 2002:2002 fetch-app:latest