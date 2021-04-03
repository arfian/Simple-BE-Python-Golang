from interfaces import ifetchrepo
import config
import requests
import json

class FetchRepo(ifetchrepo.IFetchRepo):

    def fetchResources(self):
        response = requests.get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
        return response.json()

    def fetchConversion(self):
        response = requests.get(
            "https://free.currconv.com/api/v7/convert?q=USD_IDR&compact=ultra&apiKey="+config.app['apikey'],
        )
        return json.loads(response.text)