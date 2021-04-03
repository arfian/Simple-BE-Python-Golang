from interfaces import ifetchrepo

class FetchRepo(ifetchrepo.IFetchRepo):

    def fetchResources(self):
        import requests
        response = requests.get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
        return response.text